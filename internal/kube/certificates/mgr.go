// Package certificates provides the ability to create or update
// instances of the v2alpha1 Certificate resource, and ensure that a
// corresponding Secret resource is maintained for each.
package certificates

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/skupperproject/skupper/internal/certs"
	"github.com/skupperproject/skupper/internal/kube/client"
	"github.com/skupperproject/skupper/internal/kube/watchers"
	skupperv2alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v2alpha1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// The ControllerContext interface defines the invocations the
// CertificateManager needs to make to correctly manage the resources
// it is responsible for.
type ControllerContext interface {
	// Determines whether resources in a given namespace are in
	// scope for control by the CertificateManager.
	IsControlled(namespace string) bool
	// Called to set any extra labels on resources managed by the CertificateManager.
	SetLabels(namespace string, name string, kind string, labels map[string]string) bool
	// Called to set any extra annotations on resources managed by the CertificateManager.
	SetAnnotations(namespace string, name string, kind string, annotations map[string]string) bool
}

// The CertificateManager interface defines the methods through which
// the existence of a particular Certificate resource can be
// ensured. It is currently used by package internal/kube/site.
type CertificateManager interface {
	EnsureCA(namespace string, name string, options Options) error
	Ensure(namespace string, name string, options CertOptions) error
}

type Options struct {
	Subject        string
	Refs           []metav1.OwnerReference
	ExpireInterval time.Duration
	RenewInterval  time.Duration
	Remote         bool
}

type CertOptions struct {
	Options
	CA     string
	Hosts  []string
	Client bool
	Server bool
}

type CertificateManagerImpl struct {
	definitions        map[string]*skupperv2alpha1.Certificate
	requests           map[string]*skupperv2alpha1.CertificateRequest
	secrets            map[string]*corev1.Secret
	certificateWatcher *watchers.CertificateWatcher
	requestWatcher     *watchers.CertificateRequestWatcher
	secretWatcher      *watchers.SecretWatcher
	processor          *watchers.EventProcessor
	context            ControllerContext
	logger             *slog.Logger
}

// Returns a correctly initialised CertificateManager.
func NewCertificateManager(processor *watchers.EventProcessor) *CertificateManagerImpl {
	return &CertificateManagerImpl{
		definitions: map[string]*skupperv2alpha1.Certificate{},
		secrets:     map[string]*corev1.Secret{},
		requests:    map[string]*skupperv2alpha1.CertificateRequest{},
		processor:   processor,
		logger:      slog.New(slog.Default().Handler()).With(slog.String("component", "kube.certificates.manager")),
	}
}

// Allows a ControllerContext to be set for this CertificateManager.
func (m *CertificateManagerImpl) SetControllerContext(context ControllerContext) {
	m.context = context
}

// Causes the CertificateManager to start watching relevant resources.
func (m *CertificateManagerImpl) Watch(watchNamespace string) {
	m.certificateWatcher = m.processor.WatchCertificates(watchNamespace, watchers.FilterByNamespace(m.isControlled, m.checkCertificate))
	m.requestWatcher = m.processor.WatchCertificateRequests(watchNamespace, watchers.FilterByNamespace(m.isControlled, m.checkCertificateRequest))
	m.secretWatcher = m.processor.WatchAllSecrets(watchNamespace, watchers.FilterByNamespace(m.isControlled, m.checkSecret))
}

func (m *CertificateManagerImpl) isControlled(namespace string) bool {
	if m.context != nil {
		return m.context.IsControlled(namespace)
	}
	return true
}

// This will iterate through the existing resources to recover the
// correct internal state. This should only be called after Watch()
// has been invoked.
func (m *CertificateManagerImpl) Recover() {
	for _, secret := range m.secretWatcher.List() {
		if !m.isControlled(secret.Namespace) {
			continue
		}
		m.secrets[secretKey(secret)] = secret
	}
	for _, request := range m.requestWatcher.List() {
		if !m.isControlled(request.Namespace) {
			continue
		}
		m.requests[request.Key()] = request
	}
	for _, cert := range m.certificateWatcher.List() {
		if !m.isControlled(cert.Namespace) {
			continue
		}
		if err := m.checkCertificate(cert.Key(), cert); err != nil {
			m.logger.Error("Error trying to reconcile certificate", slog.String("key", cert.Key()), slog.Any("error", err))
		}
	}
}

// This method is called to ensure that a Certificate resource exists
// to represent a CA (i.e. certificate issuer) with the properties
// specified in the arguments.
func (m *CertificateManagerImpl) EnsureCA(namespace string, name string, options Options) error {
	spec := skupperv2alpha1.CertificateSpec{
		Subject:      options.Subject,
		Signing:      true,
		RemoteIssuer: options.Remote,
	}
	if options.ExpireInterval > 0 {
		spec.ExpireInterval = options.ExpireInterval.String()
	}
	if options.RenewInterval > 0 {
		spec.RenewInterval = options.RenewInterval.String()
	}
	return m.ensure(namespace, name, spec, options.Refs)
}

// This method is called to ensure that a Certificate resource exists
// with the properties specified in the arguments. This can be called
// with different owners, in which case the owenres are all merged
// in. Hosts are tracked per owner, so if two different owners specify
// different hosts, they will all be included in the certificate, but
// if the same owner changes the hosts then they will be changed on
// the certificate. This allows the same certificate to be used for
// multiple resources such as Routes.
func (m *CertificateManagerImpl) Ensure(namespace string, name string, options CertOptions) error {
	spec := skupperv2alpha1.CertificateSpec{
		Ca:           options.CA,
		Subject:      options.Subject,
		Hosts:        options.Hosts,
		Client:       options.Client,
		Server:       options.Server,
		RemoteIssuer: options.Remote,
	}
	if options.ExpireInterval > 0 {
		spec.ExpireInterval = options.ExpireInterval.String()
	}
	if options.RenewInterval > 0 {
		spec.RenewInterval = options.RenewInterval.String()
	}
	return m.ensure(namespace, name, spec, options.Refs)
}

var compareSpecUnordered []cmp.Option = []cmp.Option{
	cmpopts.EquateEmpty(),
	cmpopts.SortSlices(func(a, b string) bool { return a < b }),
}

func (m *CertificateManagerImpl) ensure(namespace string, name string, spec skupperv2alpha1.CertificateSpec, refs []metav1.OwnerReference) error {
	key := fmt.Sprintf("%s/%s", namespace, name)
	if current, ok := m.definitions[key]; ok {
		changed := false
		ownerMap := certificateToOwnerMapping(current)
		if !ownerMap.IsControlled {
			return fmt.Errorf("certificate %q exists but is not controlled by skupper", name)
		}
		if client.MergeOwnerReferences(&current.ObjectMeta, refs) {
			changed = true
		}
		for _, ref := range refs {
			refUID := string(ref.UID)
			configuredHosts := ownerMap.PerOwnerHosts[refUID]
			if !cmp.Equal(configuredHosts, spec.Hosts, compareSpecUnordered...) {
				ownerMap.PerOwnerHosts[refUID] = spec.Hosts
			}
		}
		if ownerMap.ApplyMetadata(current) {
			changed = true
		}
		ownerRefsLength := len(current.ObjectMeta.OwnerReferences)
		if ownerRefsLength > 1 {
			// once a certificate is created and gets multiple owners ignore
			// subject changes to prevent flapping subject from differing owner
			// spec.
			spec.Subject = current.Spec.Subject
		}
		spec.Hosts = ownerMap.CombinedHosts()
		if !cmp.Equal(spec, current.Spec, compareSpecUnordered...) {
			current.Spec = spec
			changed = true
		}
		if m.context != nil {
			if current.ObjectMeta.Labels == nil {
				current.ObjectMeta.Labels = map[string]string{}
			}
			if current.ObjectMeta.Annotations == nil {
				current.ObjectMeta.Annotations = map[string]string{}
			}
			if m.context.SetLabels(namespace, name, "Certificate", current.ObjectMeta.Labels) {
				changed = true
			}
			if m.context.SetAnnotations(namespace, name, "Certificate", current.ObjectMeta.Annotations) {
				changed = true
			}
		}
		if !changed {
			return nil
		}
		updated, err := m.processor.GetSkupperClient().SkupperV2alpha1().Certificates(namespace).Update(context.Background(), current, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
		m.logger.Info("Updated certificate", slog.String("namespace", updated.Namespace), slog.String("name", updated.Name))
		m.definitions[key] = updated
		return nil
	} else {
		cert := &skupperv2alpha1.Certificate{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "skupper.io/v2alpha1",
				Kind:       "Certificate",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:            name,
				OwnerReferences: refs,
				Labels: map[string]string{
					"internal.skupper.io/certificate": "true",
				},
				Annotations: map[string]string{},
			},
			Spec: spec,
		}
		ownerMap := newOwnerMapping(refs, spec.Hosts)
		ownerMap.ApplyMetadata(cert)
		if m.context != nil {
			m.context.SetLabels(namespace, cert.Name, "Certificate", cert.ObjectMeta.Labels)
			m.context.SetAnnotations(namespace, cert.Name, "Certificate", cert.ObjectMeta.Annotations)
		}

		created, err := m.processor.GetSkupperClient().SkupperV2alpha1().Certificates(namespace).Create(context.Background(), cert, metav1.CreateOptions{})
		if err != nil {
			if !apierrors.IsAlreadyExists(err) {
				log.Printf("Error creating certificate %s: %s", key, err)
				return err
			}
			log.Printf("Certificate %s already exists - loading latest", key)
			created, err = m.processor.GetSkupperClient().SkupperV2alpha1().Certificates(namespace).Get(context.Background(), cert.Name, metav1.GetOptions{})
			if err != nil {
				return err
			}
		}
		m.definitions[key] = created
		return nil
	}
}

// Called by EventProcessor whenever there is a change to a Certificate reasource.
func (m *CertificateManagerImpl) checkCertificate(key string, certificate *skupperv2alpha1.Certificate) error {
	if certificate == nil {
		return m.certificateDeleted(key)
	}
	ownerMap := certificateToOwnerMapping(certificate)
	if ownerMap.IsControlled {
		// check for deleted owner references
		ownerUIDs := map[string]struct{}{}
		for _, ref := range certificate.OwnerReferences {
			ownerUIDs[string(ref.UID)] = struct{}{}
		}
		for configuredOwner := range ownerMap.PerOwnerHosts {
			if _, ok := ownerUIDs[configuredOwner]; !ok {
				delete(ownerMap.PerOwnerHosts, configuredOwner)
			}
		}
		if ownerMap.ApplyMetadata(certificate) {
			certificate.Spec.Hosts = ownerMap.CombinedHosts()
			updated, err := m.processor.GetSkupperClient().SkupperV2alpha1().Certificates(certificate.Namespace).Update(context.Background(), certificate, metav1.UpdateOptions{})
			if err != nil {
				return err
			}
			m.definitions[key] = updated
			return nil
		}

	}
	err := m.reconcileSecret(key, certificate, m.secrets[key])
	if err != nil {
		return err
	}
	if certificate.Spec.RemoteIssuer {
		err = m.ensureRequestFor(certificate)
	} else {
		err = m.ensureNoRequestFor(certificate)
	}
	return err
}

// This method does whatever is required to ensure that there is a
// Secret resource corresponding to the supplied CertificateResource.
func (m *CertificateManagerImpl) reconcileSecret(key string, certificate *skupperv2alpha1.Certificate, secret *corev1.Secret) error {

	var err error
	if secret != nil {
		err = m.updateSecret(key, certificate, secret)
	} else {
		err = m.createSecret(key, certificate)
	}
	return m.updateStatus(certificate, err)
}

func (m *CertificateManagerImpl) certificateDeleted(key string) error {
	var err error
	delete(m.definitions, key)
	// TODO: check if we should delete it even if not controlled
	if secret, ok := m.secrets[key]; ok {
		deleteErr := m.processor.GetKubeClient().CoreV1().Secrets(secret.Namespace).Delete(context.Background(), secret.Name, metav1.DeleteOptions{})
		if deleteErr != nil {
			m.logger.Error("Failed to delete secret after certificate removal",
				slog.String("name", secret.Name),
				slog.Any("error", deleteErr))
			err = errors.Join(err, deleteErr)
		} else {
			delete(m.secrets, key)
		}
	}
	err = errors.Join(err, m.deleteCertificateRequest(key))
	return err
}

func (m *CertificateManagerImpl) deleteCertificateRequest(key string) error {
	if certRequest, ok := m.requests[key]; ok && isResourceControlled(certRequest) {
		err := m.processor.GetSkupperClient().SkupperV2alpha1().CertificateRequests(certRequest.Namespace).Delete(context.Background(), certRequest.Name, metav1.DeleteOptions{})
		if err != nil {
			m.logger.Error("Failed to delete certificate request after certificate removal",
				slog.String("name", certRequest.Name),
				slog.Any("error", err))
			return err
		}
		delete(m.requests, key)
	}
	return nil
}

func (m *CertificateManagerImpl) secretDeleted(key string) error {
	delete(m.secrets, key)
	//TODO
	return nil
}

func (m *CertificateManagerImpl) updateStatus(certificate *skupperv2alpha1.Certificate, err error) error {
	var certRequestUpdated bool
	remoteIssuer := certificate.Spec.RemoteIssuer
	if remoteIssuer {
		if request, ok := m.requests[certificate.Key()]; ok && request.IsReady() {
			if certificate.SetReadyOrPending(true) {
				certRequestUpdated = true
			}
		} else if certificate.SetReadyOrPending(false) {
			certRequestUpdated = true
		}
	} else if certificate.SetReady(err) {
		certRequestUpdated = true
	}
	if certRequestUpdated {
		latest, err := m.processor.GetSkupperClient().SkupperV2alpha1().Certificates(certificate.Namespace).UpdateStatus(context.TODO(), certificate, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
		certificate = latest
		m.logger.Info("Updated certificate status", slog.String("namespace", certificate.Namespace), slog.String("name", certificate.Name))
		m.definitions[certificate.Key()] = latest
	} else {
		m.definitions[certificate.Key()] = certificate
	}
	return nil
}

func (m *CertificateManagerImpl) updateSecret(key string, certificate *skupperv2alpha1.Certificate, secret *corev1.Secret) error {
	changed := false
	controlled := isResourceControlled(secret)
	if !controlled {
		return errors.New("secret exists but is not controlled by skupper")
	}
	if certificate.Spec.RemoteIssuer {
		request, _ := m.requests[key]
		updSecret, err := m.createCSRSecret(certificate)
		if err != nil {
			m.logger.Error("Failed to generate updated CSR secret",
				slog.String("key", key),
				slog.Any("error", err))
			return err
		}
		if !isSecretCSRCorrect(certificate, secret) || !isRequestCSRCorrect(request, secret) {
			m.logger.Info("Certificate Request CSR updated", slog.String("key", key))
			secret.Data["tls.key"] = updSecret.Data["tls.key"]
			secret.Data["tls.csr"] = updSecret.Data["tls.csr"]
			changed = true
		} else if !isRequestCertsCorrect(key, request, secret) {
			m.logger.Info("Certificate Request certificates updated", slog.String("key", key))
			secret.Data["ca.crt"], _ = base64.StdEncoding.DecodeString(request.Status.CaCertificate)
			secret.Data["tls.crt"], _ = base64.StdEncoding.DecodeString(request.Status.Certificate)
			changed = true
		}
	} else if !isSecretCorrect(certificate, secret) {
		regenerated, err := m.generateSecret(certificate)
		if err != nil {
			m.logger.Error("Error generating Secret for Certificate",
				slog.String("namespace", certificate.Namespace),
				slog.String("name", secret.Name),
				slog.String("key", key))
			return err
		}
		changed = true
		secret.Data = regenerated.Data
	}
	if changed {
		if secret.Annotations == nil {
			secret.Annotations = map[string]string{}
		}
		secret.Annotations["internal.skupper.io/hosts"] = strings.Join(certificate.Spec.Hosts, ",")
	}
	if m.context != nil {
		if secret.Labels == nil {
			secret.Labels = map[string]string{}
		}
		if secret.Annotations == nil {
			secret.Annotations = map[string]string{}
		}
		if m.context.SetLabels(certificate.Namespace, secret.Name, "Secret", secret.Labels) {
			changed = true
		}
		if m.context.SetAnnotations(certificate.Namespace, secret.Name, "Secret", secret.Annotations) {
			changed = true
		}
	}
	if !changed {
		return nil
	}

	updated, err := m.processor.GetKubeClient().CoreV1().Secrets(certificate.Namespace).Update(context.TODO(), secret, metav1.UpdateOptions{})
	if err != nil {
		m.logger.Error("Error updating Secret for Certificate",
			slog.String("namespace", secret.Namespace),
			slog.String("name", secret.Name),
			slog.String("key", key),
			slog.Any("error", err))
		return err
	}
	m.secrets[key] = updated
	m.logger.Info("Updated Secret for Certificate",
		slog.String("namespace", secret.Namespace),
		slog.String("name", secret.Name),
		slog.String("key", key),
		slog.Any("hosts", certificate.Spec.Hosts))
	return nil
}

// isRequestCSRCorrect returns true if CertificateRequest has been defined
// and the spec.request differs from what is in the provided Secret
func isRequestCSRCorrect(request *skupperv2alpha1.CertificateRequest, secret *corev1.Secret) bool {
	if request == nil {
		return true
	}
	secretCSR, csrOk := secret.Data["tls.csr"]
	if !csrOk {
		return false
	}
	csr, _ := base64.StdEncoding.DecodeString(request.Spec.Request)
	if !bytes.Equal(secretCSR, csr) {
		return false
	}
	return true
}

// isRequestCertsCorrect returns true if CertificateRequest has a cert and a ca key populated
// and their value differ from what is in the provided Secret
func isRequestCertsCorrect(key string, request *skupperv2alpha1.CertificateRequest, secret *corev1.Secret) bool {
	if request == nil || !request.IsReady() {
		// nothing else to compare at this point
		return true
	}
	ca, _ := base64.StdEncoding.DecodeString(request.Status.CaCertificate)
	crt, _ := base64.StdEncoding.DecodeString(request.Status.Certificate)
	if len(ca) == 0 || len(crt) == 0 {
		// if no info yet available, ignore
		return true
	}
	secretCa, caOk := secret.Data["ca.crt"]
	secretCrt, crtOk := secret.Data["tls.crt"]
	if !caOk || !crtOk {
		return false
	}
	return bytes.Equal(ca, secretCa) && bytes.Equal(crt, secretCrt)
}

func (m *CertificateManagerImpl) generateSecret(certificate *skupperv2alpha1.Certificate) (*corev1.Secret, error) {
	var secret *corev1.Secret
	var err error
	if certificate.Spec.Signing {
		secret, err = certs.GenerateSecret(certificate.Name, certificate.Spec.Subject, nil, 0, nil)
		if err != nil {
			return secret, err
		}
	} else {
		expiration, err := time.ParseDuration(certificate.Spec.ExpireInterval)
		if err != nil {
			expiration = time.Hour * 24 * 365 * 5
		}
		caKey := fmt.Sprintf("%s/%s", certificate.Namespace, certificate.Spec.Ca)
		ca, ok := m.secrets[caKey]
		if !ok {
			// TODO: no CA exists yet, set error on certificate status
			return nil, fmt.Errorf("CA %q not found", caKey)
		}
		// TODO: handle server and client roles properly
		secret, err = certs.GenerateSecret(certificate.Name, certificate.Spec.Subject, certificate.Spec.Hosts, expiration, ca)
		if err != nil {
			return nil, err
		}
	}
	secret.ObjectMeta.OwnerReferences = ownerReferences(certificate)
	return secret, nil
}

func (m *CertificateManagerImpl) createSecret(key string, certificate *skupperv2alpha1.Certificate) error {
	var secret *corev1.Secret
	var err error
	if certificate.Spec.RemoteIssuer {
		secret, err = m.createCSRSecret(certificate)
	} else {
		secret, err = m.generateSecret(certificate)
	}
	if err != nil {
		m.logger.Error("Error generating secret for Certificate",
			slog.String("key", key),
			slog.Any("error", err))
		return err
	}
	secret.Annotations = map[string]string{
		"internal.skupper.io/controlled":  "true",
		"internal.skupper.io/certificate": "true",
		"internal.skupper.io/hosts":       strings.Join(certificate.Spec.Hosts, ","),
	}
	secret.Labels = map[string]string{}

	if m.context != nil {
		m.context.SetLabels(certificate.Namespace, secret.Name, "Secret", secret.Labels)
		m.context.SetAnnotations(certificate.Namespace, secret.Name, "Secret", secret.Annotations)
	}
	m.logger.Info("Creating Secret for Certificate",
		slog.String("namespace", certificate.Namespace),
		slog.String("name", secret.Name),
		slog.String("key", key),
		slog.Any("hosts", certificate.Spec.Hosts))
	created, err := m.processor.GetKubeClient().CoreV1().Secrets(certificate.Namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		m.logger.Error("Error creating Secret for Certificate",
			slog.String("namespace", certificate.Namespace),
			slog.String("name", secret.Name),
			slog.String("key", key),
			slog.Any("error", err))
		return err
	}
	m.secrets[key] = created
	m.logger.Info("Created Secret for Certificate",
		slog.String("namespace", certificate.Namespace),
		slog.String("name", secret.Name),
		slog.String("key", key),
		slog.Any("hosts", certificate.Spec.Hosts))
	return nil
}

// Called by EventProcessor whenever there is a change in a relevant
// Secret resource.
func (m *CertificateManagerImpl) checkSecret(key string, secret *corev1.Secret) error {
	if secret == nil {
		return m.secretDeleted(key)
	}
	m.secrets[key] = secret
	if definition, ok := m.definitions[key]; ok {
		return m.reconcileSecret(key, definition, secret)
	}

	return nil
}

func (m *CertificateManagerImpl) createCSRSecret(certificate *skupperv2alpha1.Certificate) (*corev1.Secret, error) {
	return certs.GenerateCSRSecret(certificate.Name, certificate.Spec.Subject, certificate.Spec.Hosts, certificate.Spec.Signing)
}

func (m *CertificateManagerImpl) ensureRequestFor(certificate *skupperv2alpha1.Certificate) error {
	crClient := m.processor.GetSkupperClient().SkupperV2alpha1().CertificateRequests(certificate.Namespace)
	if existing, ok := m.requests[certificate.Key()]; ok {
		changed := false
		csrStr, _ := base64.StdEncoding.DecodeString(existing.Spec.Request)
		secret, ok := m.secrets[certificate.Key()]
		if !ok {
			return nil
		}
		secretCsrStr, ok := secret.Data["tls.csr"]
		if ok && !bytes.Equal(csrStr, secretCsrStr) {
			existing.Spec.Request = base64.StdEncoding.EncodeToString(secretCsrStr)
			changed = true
		}
		if certificate.Spec.Ca != existing.Spec.Issuer {
			existing.Spec.Issuer = certificate.Spec.Ca
			changed = true
		}
		if !changed {
			return nil
		}
		m.logger.Info("Updating CertificateRequest", slog.String("key", certificate.Key()))
		updated, err := crClient.Update(context.Background(), existing, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
		m.requests[certificate.Key()] = updated
	} else {
		secret, secretFound := m.secrets[certificate.Key()]
		if !secretFound {
			m.logger.Info("Secret does not exist, CertificateRequest cannot be created yet",
				slog.String("key", certificate.Key()))
			return nil
		}
		csrStr, csrFound := secret.Data["tls.csr"]
		if !csrFound {
			m.logger.Info("Secret does not contain a CSR, CertificateRequest cannot be created yet",
				slog.String("key", certificate.Key()))
			return nil
		}
		newRequest := &skupperv2alpha1.CertificateRequest{
			TypeMeta: metav1.TypeMeta{
				APIVersion: "skupper.io/v2alpha1",
				Kind:       "CertificateRequest",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:            certificate.Name,
				OwnerReferences: ownerReferences(certificate),
				Annotations: map[string]string{
					"internal.skupper.io/controlled": "true",
				},
			},
			Spec: skupperv2alpha1.CertificateRequestSpec{
				Issuer:  certificate.Spec.Ca,
				Request: base64.StdEncoding.EncodeToString(csrStr),
			},
		}
		m.logger.Info("Creating CertificateRequest", slog.String("key", certificate.Key()))
		created, err := crClient.Create(context.Background(), newRequest, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		m.requests[certificate.Key()] = created
	}
	return nil
}

func isSecretCorrect(certificate *skupperv2alpha1.Certificate, secret *corev1.Secret) bool {
	data, ok := secret.Data["tls.crt"]
	if !ok {
		return false
	}
	cert, err := certs.DecodeCertificate(data)
	if err != nil {
		slog.Error("Bad certificate secret", slog.String("key", certificate.Key()), slog.Any("error", err))
		return false
	}
	if time.Now().After(cert.NotAfter) {
		slog.Info("Certificate has expired", slog.String("key", certificate.Key()))
		return false
	}
	if certificate.Spec.RenewInterval != "" {
		if renewInterval, err := time.ParseDuration(certificate.Spec.RenewInterval); err == nil {
			renewTime := cert.NotBefore.Add(renewInterval)
			if time.Now().After(renewTime) {
				slog.Info("Certificate will be renewed now",
					slog.String("key", certificate.Key()),
					slog.Any("renewInterval", renewInterval),
					slog.Any("after", cert.NotBefore))
				return false
			}
		}
	}
	if certificate.Spec.Subject != cert.Subject.CommonName {
		return false
	}
	return isCertificateHostsCorrect(certificate, cert.DNSNames, cert.IPAddresses)
}

func isCertificateHostsCorrect(certificate *skupperv2alpha1.Certificate, dnsNames []string, ipAddresses []net.IP) bool {
	validFor := map[string]string{}
	for _, host := range dnsNames {
		// Ignore empty DNSNames - GH-2277
		if host == "" {
			continue
		}
		validFor[host] = host
	}
	for _, ip := range ipAddresses {
		validFor[ip.String()] = ip.String()
	}
	if len(certificate.Spec.Hosts) != len(validFor) {
		return false
	}
	for _, host := range certificate.Spec.Hosts {
		if _, ok := validFor[host]; !ok {
			return false
		}
	}
	return true
}

func isSecretCSRCorrect(certificate *skupperv2alpha1.Certificate, secret *corev1.Secret) bool {
	data, ok := secret.Data["tls.csr"]
	if !ok {
		return false
	}
	request, err := certs.DecodeCSR(data)
	if err != nil {
		slog.Info("Bad certificate request secret",
			slog.String("key", certificate.Key()),
			slog.Any("error", err))
		return false
	}
	if certificate.Spec.Subject != request.Subject.CommonName {
		return false
	}
	return isCertificateHostsCorrect(certificate, request.DNSNames, request.IPAddresses)
}

func isResourceControlled(obj metav1.ObjectMetaAccessor) bool {
	return hasControlledAnnotation(obj.GetObjectMeta()) || hasCertificateOwner(obj.GetObjectMeta())
}

func hasControlledAnnotation(metadata metav1.Object) bool {
	if metadata.GetAnnotations() == nil {
		return false
	}
	_, ok := metadata.GetAnnotations()["internal.skupper.io/controlled"]
	return ok
}

func hasCertificateOwner(metadata metav1.Object) bool {
	for _, owner := range metadata.GetOwnerReferences() {
		if owner.Kind == "Certificate" && owner.APIVersion == "skupper.io/v2alpha1" {
			return true
		}
	}
	return false
}

func secretKey(secret *corev1.Secret) string {
	return fmt.Sprintf("%s/%s", secret.Namespace, secret.Name)
}

func ownerReferences(cert *skupperv2alpha1.Certificate) []metav1.OwnerReference {
	return []metav1.OwnerReference{
		{
			Kind:       "Certificate",
			APIVersion: "skupper.io/v2alpha1",
			Name:       cert.Name,
			UID:        cert.ObjectMeta.UID,
		},
	}
}

func (m *CertificateManagerImpl) checkCertificateRequest(key string, request *skupperv2alpha1.CertificateRequest) error {
	if request == nil {
		delete(m.requests, key)
		return nil
	}
	m.requests[key] = request
	return nil
}

func (m *CertificateManagerImpl) ensureNoRequestFor(certificate *skupperv2alpha1.Certificate) error {
	var err error
	if m.requests[certificate.Key()] != nil {
		log.Printf("Deleting CertificateRequest and Secret for %s (remoteIssuer disabled)", certificate.Key())
		secretClient := m.processor.GetKubeClient().CoreV1().Secrets(certificate.Namespace)
		err = secretClient.Delete(context.Background(), certificate.Name, metav1.DeleteOptions{})
		if err != nil {
			return err
		}
		_ = m.secretDeleted(certificate.Key())
		err = m.deleteCertificateRequest(certificate.Key())
	}
	return err
}
