/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package certs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func publicKey(priv interface{}) interface{} {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	default:
		return nil
	}
}

func pemBlockForKey(priv interface{}) *pem.Block {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	default:
		return nil
	}
}

type CertificateAuthority struct {
	Certificate *x509.Certificate
	Key         interface{}
	CrtData     []byte
}

type CertificateData map[string][]byte

func decodeDataElement(in []byte, name string) ([]byte, error) {
	block, _ := pem.Decode(in)
	if block == nil {
		return nil, fmt.Errorf("failed to read PEM encoded data from %q", name)
	}
	return block.Bytes, nil
}

func getCAFromSecret(secret *corev1.Secret) (*CertificateAuthority, error) {
	if secret == nil || secret.Data == nil {
		return nil, nil
	}

	certBytes, err := decodeDataElement(secret.Data["tls.crt"], "tls.crt")
	if err != nil {
		return nil, err
	}

	cert, err := x509.ParseCertificate(certBytes)
	if err != nil {
		return nil, err
	}

	privateKeyBytes, err := decodeDataElement(secret.Data["tls.key"], "tls.key")
	if err != nil {
		return nil, err
	}
	key, err := x509.ParsePKCS1PrivateKey(privateKeyBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to get CA private key from secret %s", err)
	}

	return &CertificateAuthority{
		Certificate: cert,
		Key:         key,
		CrtData:     secret.Data["tls.crt"],
	}, nil
}

// GenerateSecret generates a kubernetes secret.
// name is the corev1.Secret's name.
// subject is the x509 certificate's common name.
// hosts are the host names in the x509 certificate subject alternative names extension.
// expiration is when the secret expires, if zero is passed in, the expiration is set to 5 years from now
// ca is the certificate authority, if nil a ca cert will be created.
func GenerateSecret(name string, subject string, hosts []string, expiration time.Duration, ca *corev1.Secret) (*corev1.Secret, error) {
	caCert, err := getCAFromSecret(ca)

	if err != nil {
		return nil, fmt.Errorf("error reading CA Certificate from Secret %q: %s", ca.Name, err)
	}

	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)

	}

	notBefore := time.Now()
	if expiration == 0 {
		expiration = 5 * 365 * 24 * time.Hour
	}
	notAfter := notBefore.Add(expiration)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			CommonName: subject,
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	for _, h := range hosts {
		// Remove leading and trailing whitespaces from the string
		h = strings.TrimSpace(h)
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		}
		template.DNSNames = append(template.DNSNames, h)
	}

	var parent *x509.Certificate
	var cakey interface{}
	if caCert == nil {
		// self signed
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
		parent = &template
		cakey = priv
	} else {
		parent = caCert.Certificate
		cakey = caCert.Key
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, parent, publicKey(priv), cakey)
	if err != nil {
		return nil, err
	}

	secret := corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Type: "kubernetes.io/tls",
		Data: map[string][]byte{},
	}

	certString := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	keyString := pem.EncodeToMemory(pemBlockForKey(priv))

	secret.Data["tls.crt"] = []byte(certString)
	secret.Data["tls.key"] = []byte(keyString)
	if ca != nil {
		secret.Data["ca.crt"] = caCert.CrtData
	} else {
		secret.Data["ca.crt"] = secret.Data["tls.crt"] //self.signed
	}

	return &secret, nil
}

func DecodeCertificate(data []byte) (*x509.Certificate, error) {
	b, _ := pem.Decode(data)
	if b == nil {
		return nil, fmt.Errorf("Could not decode PEM block from data")
	}
	return x509.ParseCertificate(b.Bytes)
}

// Define the BasicConstraints ASN.1 sequence
type basicConstraints struct {
	IsCA       bool `asn1:"optional"`
	MaxPathLen int  `asn1:"optional,default:-1"`
}

func GenerateCSRSecret(name string, subject string, hosts []string, signing bool) (*corev1.Secret, error) {
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %v", err)

	}

	template := x509.CertificateRequest{
		PublicKey: publicKey(priv),
		Subject: pkix.Name{
			CommonName: subject,
		},
	}

	for _, h := range hosts {
		// Remove leading and trailing whitespaces from the string
		h = strings.TrimSpace(h)
		if ip := net.ParseIP(h); ip != nil {
			template.IPAddresses = append(template.IPAddresses, ip)
		}
		template.DNSNames = append(template.DNSNames, h)
	}

	// Need to go through RFC 5280, see:
	// https://oid-info.com
	// https://oidref.com
	oidExtensionKeyUsage := asn1.ObjectIdentifier{2, 5, 29, 15}
	kuValue, _ := asn1.Marshal(asn1.BitString{Bytes: []byte{byte(x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature)}, BitLength: 3})

	oidExtensionExtendedKeyUsage := asn1.ObjectIdentifier{2, 5, 29, 37}
	oidExtKeyUsageServerAuth := asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 1}
	oidExtKeyUsageClientAuth := asn1.ObjectIdentifier{1, 3, 6, 1, 5, 5, 7, 3, 2}
	ekuDER, err := asn1.Marshal([]asn1.ObjectIdentifier{oidExtKeyUsageServerAuth, oidExtKeyUsageClientAuth})
	if err != nil {
		return nil, err
	}
	template.ExtraExtensions = []pkix.Extension{
		{
			Id:       oidExtensionKeyUsage,
			Critical: true,
			Value:    kuValue,
		},
		{
			Id:       oidExtensionExtendedKeyUsage,
			Critical: false,
			Value:    ekuDER,
		},
	}
	if signing {
		caVal, _ := asn1.Marshal(basicConstraints{IsCA: true, MaxPathLen: 0})
		caExtension := pkix.Extension{
			Id:       asn1.ObjectIdentifier{2, 5, 29, 19},
			Critical: true,
			Value:    caVal,
		}
		template.ExtraExtensions = append(template.ExtraExtensions, caExtension)
	}

	derBytes, err := x509.CreateCertificateRequest(rand.Reader, &template, priv)
	if err != nil {
		return nil, err
	}

	secret := corev1.Secret{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Secret",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Type: "kubernetes.io/tls",
		Data: map[string][]byte{},
	}

	csrString := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: derBytes})
	keyString := pem.EncodeToMemory(pemBlockForKey(priv))

	secret.Data["tls.csr"] = csrString
	secret.Data["tls.key"] = keyString
	secret.Data["tls.crt"] = []byte{}

	return &secret, nil
}

func DecodeCSR(data []byte) (*x509.CertificateRequest, error) {
	b, _ := pem.Decode(data)
	if b == nil {
		return nil, fmt.Errorf("Could not decode PEM block from data")
	}
	return x509.ParseCertificateRequest(b.Bytes)
}
