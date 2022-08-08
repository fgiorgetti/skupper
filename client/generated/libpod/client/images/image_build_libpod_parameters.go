// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewImageBuildLibpodParams creates a new ImageBuildLibpodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageBuildLibpodParams() *ImageBuildLibpodParams {
	return &ImageBuildLibpodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageBuildLibpodParamsWithTimeout creates a new ImageBuildLibpodParams object
// with the ability to set a timeout on a request.
func NewImageBuildLibpodParamsWithTimeout(timeout time.Duration) *ImageBuildLibpodParams {
	return &ImageBuildLibpodParams{
		timeout: timeout,
	}
}

// NewImageBuildLibpodParamsWithContext creates a new ImageBuildLibpodParams object
// with the ability to set a context for a request.
func NewImageBuildLibpodParamsWithContext(ctx context.Context) *ImageBuildLibpodParams {
	return &ImageBuildLibpodParams{
		Context: ctx,
	}
}

// NewImageBuildLibpodParamsWithHTTPClient creates a new ImageBuildLibpodParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageBuildLibpodParamsWithHTTPClient(client *http.Client) *ImageBuildLibpodParams {
	return &ImageBuildLibpodParams{
		HTTPClient: client,
	}
}

/* ImageBuildLibpodParams contains all the parameters to send to the API endpoint
   for the image build libpod operation.

   Typically these are written to a http.Request.
*/
type ImageBuildLibpodParams struct {

	/* Allplatforms.

	     Instead of building for a set of platforms specified using the platform option, inspect the build's base images,
	and build for all of the platforms that are available.  Stages that use *scratch* as a starting point can not be inspected,
	so at least one non-*scratch* stage must be present for detection to work usefully.

	*/
	Allplatforms *bool

	/* Buildargs.

	     JSON map of string pairs denoting build-time variables.
	For example, the build argument `Foo` with the value of `bar` would be encoded in JSON as `["Foo":"bar"]`.

	For example, buildargs={"Foo":"bar"}.

	Note(s):
	* This should not be used to pass secrets.
	* The value of buildargs should be URI component encoded before being passed to the API.

	(As of version 1.xx)

	*/
	Buildargs *string

	/* Cachefrom.

	     JSON array of images used to build cache resolution
	(As of version 1.xx)

	*/
	Cachefrom *string

	/* Cpuperiod.

	     CPUPeriod limits the CPU CFS (Completely Fair Scheduler) period
	(As of version 1.xx)

	*/
	Cpuperiod *int64

	/* Cpuquota.

	     CPUQuota limits the CPU CFS (Completely Fair Scheduler) quota
	(As of version 1.xx)

	*/
	Cpuquota *int64

	/* Cpusetcpus.

	     CPUSetCPUs in which to allow execution (0-3, 0,1)
	(As of version 1.xx)

	*/
	Cpusetcpus *string

	/* Cpushares.

	     CPUShares (relative weight
	(As of version 1.xx)

	*/
	Cpushares *int64

	/* Dockerfile.

	     Path within the build context to the `Dockerfile`.
	This is ignored if remote is specified and points to an external `Dockerfile`.


	     Default: "Dockerfile"
	*/
	Dockerfile *string

	/* Extrahosts.

	     TBD Extra hosts to add to /etc/hosts
	(As of version 1.xx)

	*/
	Extrahosts *string

	/* Forcerm.

	     Always remove intermediate containers, even upon failure
	(As of version 1.xx)

	*/
	Forcerm *bool

	/* Httpproxy.

	     Inject http proxy environment variables into container
	(As of version 2.0.0)

	*/
	Httpproxy *bool

	/* Labels.

	     JSON map of key, value pairs to set as labels on the new image
	(As of version 1.xx)

	*/
	Labels *string

	/* Layers.

	     Cache intermediate layers during build.
	(As of version 1.xx)


	     Default: true
	*/
	Layers *bool

	/* Memory.

	     Memory is the upper limit (in bytes) on how much memory running containers can use
	(As of version 1.xx)

	*/
	Memory *int64

	/* Memswap.

	     MemorySwap limits the amount of memory and swap together
	(As of version 1.xx)

	*/
	Memswap *int64

	/* Networkmode.

	     Sets the networking mode for the run commands during build.
	Supported standard values are:
	  * `bridge` limited to containers within a single host, port mapping required for external access
	  * `host` no isolation between host and containers on this network
	  * `none` disable all networking for this container
	  * container:<nameOrID> share networking with given container
	  ---All other values are assumed to be a custom network's name
	(As of version 1.xx)


	     Default: "bridge"
	*/
	Networkmode *string

	/* Nocache.

	     Do not use the cache when building the image
	(As of version 1.xx)

	*/
	Nocache *bool

	/* Outputs.

	     output configuration TBD
	(As of version 1.xx)

	*/
	Outputs *string

	/* Platform.

	     Platform format os[/arch[/variant]]
	(As of version 1.xx)

	*/
	Platform *string

	/* Pull.

	     Attempt to pull the image even if an older image exists locally
	(As of version 1.xx)

	*/
	Pull *bool

	/* Q.

	   Suppress verbose build output

	*/
	Q *bool

	/* Remote.

	     A Git repository URI or HTTP/HTTPS context URI.
	If the URI points to a single text file, the file’s contents are placed
	into a file called Dockerfile and the image is built from that file. If
	the URI points to a tarball, the file is downloaded by the daemon and the
	contents therein used as the context for the build. If the URI points to a
	tarball and the dockerfile parameter is also specified, there must be a file
	with the corresponding path inside the tarball.
	(As of version 1.xx)

	*/
	Remote *string

	/* Rm.

	     Remove intermediate containers after a successful build
	(As of version 1.xx)


	     Default: true
	*/
	Rm *bool

	/* Shmsize.

	     ShmSize is the "size" value to use when mounting an shmfs on the container's /dev/shm directory.
	Default is 64MB
	(As of version 1.xx)


	     Default: 67108864
	*/
	Shmsize *int64

	/* Squash.

	     Silently ignored.
	Squash the resulting images layers into a single layer
	(As of version 1.xx)

	*/
	Squash *bool

	/* T.

	   A name and optional tag to apply to the image in the `name:tag` format.  If you omit the tag the default latest value is assumed. You can provide several t parameters.

	   Default: "latest"
	*/
	T *string

	/* Target.

	     Target build stage
	(As of version 1.xx)

	*/
	Target *string

	/* Unsetenv.

	   Unset environment variables from the final image.
	*/
	Unsetenv []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image build libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageBuildLibpodParams) WithDefaults() *ImageBuildLibpodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image build libpod params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageBuildLibpodParams) SetDefaults() {
	var (
		allplatformsDefault = bool(false)

		dockerfileDefault = string("Dockerfile")

		forcermDefault = bool(false)

		layersDefault = bool(true)

		networkmodeDefault = string("bridge")

		nocacheDefault = bool(false)

		pullDefault = bool(false)

		qDefault = bool(false)

		rmDefault = bool(true)

		shmsizeDefault = int64(6.7108864e+07)

		squashDefault = bool(false)

		tDefault = string("latest")
	)

	val := ImageBuildLibpodParams{
		Allplatforms: &allplatformsDefault,
		Dockerfile:   &dockerfileDefault,
		Forcerm:      &forcermDefault,
		Layers:       &layersDefault,
		Networkmode:  &networkmodeDefault,
		Nocache:      &nocacheDefault,
		Pull:         &pullDefault,
		Q:            &qDefault,
		Rm:           &rmDefault,
		Shmsize:      &shmsizeDefault,
		Squash:       &squashDefault,
		T:            &tDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the image build libpod params
func (o *ImageBuildLibpodParams) WithTimeout(timeout time.Duration) *ImageBuildLibpodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image build libpod params
func (o *ImageBuildLibpodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image build libpod params
func (o *ImageBuildLibpodParams) WithContext(ctx context.Context) *ImageBuildLibpodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image build libpod params
func (o *ImageBuildLibpodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image build libpod params
func (o *ImageBuildLibpodParams) WithHTTPClient(client *http.Client) *ImageBuildLibpodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image build libpod params
func (o *ImageBuildLibpodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAllplatforms adds the allplatforms to the image build libpod params
func (o *ImageBuildLibpodParams) WithAllplatforms(allplatforms *bool) *ImageBuildLibpodParams {
	o.SetAllplatforms(allplatforms)
	return o
}

// SetAllplatforms adds the allplatforms to the image build libpod params
func (o *ImageBuildLibpodParams) SetAllplatforms(allplatforms *bool) {
	o.Allplatforms = allplatforms
}

// WithBuildargs adds the buildargs to the image build libpod params
func (o *ImageBuildLibpodParams) WithBuildargs(buildargs *string) *ImageBuildLibpodParams {
	o.SetBuildargs(buildargs)
	return o
}

// SetBuildargs adds the buildargs to the image build libpod params
func (o *ImageBuildLibpodParams) SetBuildargs(buildargs *string) {
	o.Buildargs = buildargs
}

// WithCachefrom adds the cachefrom to the image build libpod params
func (o *ImageBuildLibpodParams) WithCachefrom(cachefrom *string) *ImageBuildLibpodParams {
	o.SetCachefrom(cachefrom)
	return o
}

// SetCachefrom adds the cachefrom to the image build libpod params
func (o *ImageBuildLibpodParams) SetCachefrom(cachefrom *string) {
	o.Cachefrom = cachefrom
}

// WithCpuperiod adds the cpuperiod to the image build libpod params
func (o *ImageBuildLibpodParams) WithCpuperiod(cpuperiod *int64) *ImageBuildLibpodParams {
	o.SetCpuperiod(cpuperiod)
	return o
}

// SetCpuperiod adds the cpuperiod to the image build libpod params
func (o *ImageBuildLibpodParams) SetCpuperiod(cpuperiod *int64) {
	o.Cpuperiod = cpuperiod
}

// WithCpuquota adds the cpuquota to the image build libpod params
func (o *ImageBuildLibpodParams) WithCpuquota(cpuquota *int64) *ImageBuildLibpodParams {
	o.SetCpuquota(cpuquota)
	return o
}

// SetCpuquota adds the cpuquota to the image build libpod params
func (o *ImageBuildLibpodParams) SetCpuquota(cpuquota *int64) {
	o.Cpuquota = cpuquota
}

// WithCpusetcpus adds the cpusetcpus to the image build libpod params
func (o *ImageBuildLibpodParams) WithCpusetcpus(cpusetcpus *string) *ImageBuildLibpodParams {
	o.SetCpusetcpus(cpusetcpus)
	return o
}

// SetCpusetcpus adds the cpusetcpus to the image build libpod params
func (o *ImageBuildLibpodParams) SetCpusetcpus(cpusetcpus *string) {
	o.Cpusetcpus = cpusetcpus
}

// WithCpushares adds the cpushares to the image build libpod params
func (o *ImageBuildLibpodParams) WithCpushares(cpushares *int64) *ImageBuildLibpodParams {
	o.SetCpushares(cpushares)
	return o
}

// SetCpushares adds the cpushares to the image build libpod params
func (o *ImageBuildLibpodParams) SetCpushares(cpushares *int64) {
	o.Cpushares = cpushares
}

// WithDockerfile adds the dockerfile to the image build libpod params
func (o *ImageBuildLibpodParams) WithDockerfile(dockerfile *string) *ImageBuildLibpodParams {
	o.SetDockerfile(dockerfile)
	return o
}

// SetDockerfile adds the dockerfile to the image build libpod params
func (o *ImageBuildLibpodParams) SetDockerfile(dockerfile *string) {
	o.Dockerfile = dockerfile
}

// WithExtrahosts adds the extrahosts to the image build libpod params
func (o *ImageBuildLibpodParams) WithExtrahosts(extrahosts *string) *ImageBuildLibpodParams {
	o.SetExtrahosts(extrahosts)
	return o
}

// SetExtrahosts adds the extrahosts to the image build libpod params
func (o *ImageBuildLibpodParams) SetExtrahosts(extrahosts *string) {
	o.Extrahosts = extrahosts
}

// WithForcerm adds the forcerm to the image build libpod params
func (o *ImageBuildLibpodParams) WithForcerm(forcerm *bool) *ImageBuildLibpodParams {
	o.SetForcerm(forcerm)
	return o
}

// SetForcerm adds the forcerm to the image build libpod params
func (o *ImageBuildLibpodParams) SetForcerm(forcerm *bool) {
	o.Forcerm = forcerm
}

// WithHttpproxy adds the httpproxy to the image build libpod params
func (o *ImageBuildLibpodParams) WithHttpproxy(httpproxy *bool) *ImageBuildLibpodParams {
	o.SetHttpproxy(httpproxy)
	return o
}

// SetHttpproxy adds the httpproxy to the image build libpod params
func (o *ImageBuildLibpodParams) SetHttpproxy(httpproxy *bool) {
	o.Httpproxy = httpproxy
}

// WithLabels adds the labels to the image build libpod params
func (o *ImageBuildLibpodParams) WithLabels(labels *string) *ImageBuildLibpodParams {
	o.SetLabels(labels)
	return o
}

// SetLabels adds the labels to the image build libpod params
func (o *ImageBuildLibpodParams) SetLabels(labels *string) {
	o.Labels = labels
}

// WithLayers adds the layers to the image build libpod params
func (o *ImageBuildLibpodParams) WithLayers(layers *bool) *ImageBuildLibpodParams {
	o.SetLayers(layers)
	return o
}

// SetLayers adds the layers to the image build libpod params
func (o *ImageBuildLibpodParams) SetLayers(layers *bool) {
	o.Layers = layers
}

// WithMemory adds the memory to the image build libpod params
func (o *ImageBuildLibpodParams) WithMemory(memory *int64) *ImageBuildLibpodParams {
	o.SetMemory(memory)
	return o
}

// SetMemory adds the memory to the image build libpod params
func (o *ImageBuildLibpodParams) SetMemory(memory *int64) {
	o.Memory = memory
}

// WithMemswap adds the memswap to the image build libpod params
func (o *ImageBuildLibpodParams) WithMemswap(memswap *int64) *ImageBuildLibpodParams {
	o.SetMemswap(memswap)
	return o
}

// SetMemswap adds the memswap to the image build libpod params
func (o *ImageBuildLibpodParams) SetMemswap(memswap *int64) {
	o.Memswap = memswap
}

// WithNetworkmode adds the networkmode to the image build libpod params
func (o *ImageBuildLibpodParams) WithNetworkmode(networkmode *string) *ImageBuildLibpodParams {
	o.SetNetworkmode(networkmode)
	return o
}

// SetNetworkmode adds the networkmode to the image build libpod params
func (o *ImageBuildLibpodParams) SetNetworkmode(networkmode *string) {
	o.Networkmode = networkmode
}

// WithNocache adds the nocache to the image build libpod params
func (o *ImageBuildLibpodParams) WithNocache(nocache *bool) *ImageBuildLibpodParams {
	o.SetNocache(nocache)
	return o
}

// SetNocache adds the nocache to the image build libpod params
func (o *ImageBuildLibpodParams) SetNocache(nocache *bool) {
	o.Nocache = nocache
}

// WithOutputs adds the outputs to the image build libpod params
func (o *ImageBuildLibpodParams) WithOutputs(outputs *string) *ImageBuildLibpodParams {
	o.SetOutputs(outputs)
	return o
}

// SetOutputs adds the outputs to the image build libpod params
func (o *ImageBuildLibpodParams) SetOutputs(outputs *string) {
	o.Outputs = outputs
}

// WithPlatform adds the platform to the image build libpod params
func (o *ImageBuildLibpodParams) WithPlatform(platform *string) *ImageBuildLibpodParams {
	o.SetPlatform(platform)
	return o
}

// SetPlatform adds the platform to the image build libpod params
func (o *ImageBuildLibpodParams) SetPlatform(platform *string) {
	o.Platform = platform
}

// WithPull adds the pull to the image build libpod params
func (o *ImageBuildLibpodParams) WithPull(pull *bool) *ImageBuildLibpodParams {
	o.SetPull(pull)
	return o
}

// SetPull adds the pull to the image build libpod params
func (o *ImageBuildLibpodParams) SetPull(pull *bool) {
	o.Pull = pull
}

// WithQ adds the q to the image build libpod params
func (o *ImageBuildLibpodParams) WithQ(q *bool) *ImageBuildLibpodParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the image build libpod params
func (o *ImageBuildLibpodParams) SetQ(q *bool) {
	o.Q = q
}

// WithRemote adds the remote to the image build libpod params
func (o *ImageBuildLibpodParams) WithRemote(remote *string) *ImageBuildLibpodParams {
	o.SetRemote(remote)
	return o
}

// SetRemote adds the remote to the image build libpod params
func (o *ImageBuildLibpodParams) SetRemote(remote *string) {
	o.Remote = remote
}

// WithRm adds the rm to the image build libpod params
func (o *ImageBuildLibpodParams) WithRm(rm *bool) *ImageBuildLibpodParams {
	o.SetRm(rm)
	return o
}

// SetRm adds the rm to the image build libpod params
func (o *ImageBuildLibpodParams) SetRm(rm *bool) {
	o.Rm = rm
}

// WithShmsize adds the shmsize to the image build libpod params
func (o *ImageBuildLibpodParams) WithShmsize(shmsize *int64) *ImageBuildLibpodParams {
	o.SetShmsize(shmsize)
	return o
}

// SetShmsize adds the shmsize to the image build libpod params
func (o *ImageBuildLibpodParams) SetShmsize(shmsize *int64) {
	o.Shmsize = shmsize
}

// WithSquash adds the squash to the image build libpod params
func (o *ImageBuildLibpodParams) WithSquash(squash *bool) *ImageBuildLibpodParams {
	o.SetSquash(squash)
	return o
}

// SetSquash adds the squash to the image build libpod params
func (o *ImageBuildLibpodParams) SetSquash(squash *bool) {
	o.Squash = squash
}

// WithT adds the t to the image build libpod params
func (o *ImageBuildLibpodParams) WithT(t *string) *ImageBuildLibpodParams {
	o.SetT(t)
	return o
}

// SetT adds the t to the image build libpod params
func (o *ImageBuildLibpodParams) SetT(t *string) {
	o.T = t
}

// WithTarget adds the target to the image build libpod params
func (o *ImageBuildLibpodParams) WithTarget(target *string) *ImageBuildLibpodParams {
	o.SetTarget(target)
	return o
}

// SetTarget adds the target to the image build libpod params
func (o *ImageBuildLibpodParams) SetTarget(target *string) {
	o.Target = target
}

// WithUnsetenv adds the unsetenv to the image build libpod params
func (o *ImageBuildLibpodParams) WithUnsetenv(unsetenv []string) *ImageBuildLibpodParams {
	o.SetUnsetenv(unsetenv)
	return o
}

// SetUnsetenv adds the unsetenv to the image build libpod params
func (o *ImageBuildLibpodParams) SetUnsetenv(unsetenv []string) {
	o.Unsetenv = unsetenv
}

// WriteToRequest writes these params to a swagger request
func (o *ImageBuildLibpodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Allplatforms != nil {

		// query param allplatforms
		var qrAllplatforms bool

		if o.Allplatforms != nil {
			qrAllplatforms = *o.Allplatforms
		}
		qAllplatforms := swag.FormatBool(qrAllplatforms)
		if qAllplatforms != "" {

			if err := r.SetQueryParam("allplatforms", qAllplatforms); err != nil {
				return err
			}
		}
	}

	if o.Buildargs != nil {

		// query param buildargs
		var qrBuildargs string

		if o.Buildargs != nil {
			qrBuildargs = *o.Buildargs
		}
		qBuildargs := qrBuildargs
		if qBuildargs != "" {

			if err := r.SetQueryParam("buildargs", qBuildargs); err != nil {
				return err
			}
		}
	}

	if o.Cachefrom != nil {

		// query param cachefrom
		var qrCachefrom string

		if o.Cachefrom != nil {
			qrCachefrom = *o.Cachefrom
		}
		qCachefrom := qrCachefrom
		if qCachefrom != "" {

			if err := r.SetQueryParam("cachefrom", qCachefrom); err != nil {
				return err
			}
		}
	}

	if o.Cpuperiod != nil {

		// query param cpuperiod
		var qrCpuperiod int64

		if o.Cpuperiod != nil {
			qrCpuperiod = *o.Cpuperiod
		}
		qCpuperiod := swag.FormatInt64(qrCpuperiod)
		if qCpuperiod != "" {

			if err := r.SetQueryParam("cpuperiod", qCpuperiod); err != nil {
				return err
			}
		}
	}

	if o.Cpuquota != nil {

		// query param cpuquota
		var qrCpuquota int64

		if o.Cpuquota != nil {
			qrCpuquota = *o.Cpuquota
		}
		qCpuquota := swag.FormatInt64(qrCpuquota)
		if qCpuquota != "" {

			if err := r.SetQueryParam("cpuquota", qCpuquota); err != nil {
				return err
			}
		}
	}

	if o.Cpusetcpus != nil {

		// query param cpusetcpus
		var qrCpusetcpus string

		if o.Cpusetcpus != nil {
			qrCpusetcpus = *o.Cpusetcpus
		}
		qCpusetcpus := qrCpusetcpus
		if qCpusetcpus != "" {

			if err := r.SetQueryParam("cpusetcpus", qCpusetcpus); err != nil {
				return err
			}
		}
	}

	if o.Cpushares != nil {

		// query param cpushares
		var qrCpushares int64

		if o.Cpushares != nil {
			qrCpushares = *o.Cpushares
		}
		qCpushares := swag.FormatInt64(qrCpushares)
		if qCpushares != "" {

			if err := r.SetQueryParam("cpushares", qCpushares); err != nil {
				return err
			}
		}
	}

	if o.Dockerfile != nil {

		// query param dockerfile
		var qrDockerfile string

		if o.Dockerfile != nil {
			qrDockerfile = *o.Dockerfile
		}
		qDockerfile := qrDockerfile
		if qDockerfile != "" {

			if err := r.SetQueryParam("dockerfile", qDockerfile); err != nil {
				return err
			}
		}
	}

	if o.Extrahosts != nil {

		// query param extrahosts
		var qrExtrahosts string

		if o.Extrahosts != nil {
			qrExtrahosts = *o.Extrahosts
		}
		qExtrahosts := qrExtrahosts
		if qExtrahosts != "" {

			if err := r.SetQueryParam("extrahosts", qExtrahosts); err != nil {
				return err
			}
		}
	}

	if o.Forcerm != nil {

		// query param forcerm
		var qrForcerm bool

		if o.Forcerm != nil {
			qrForcerm = *o.Forcerm
		}
		qForcerm := swag.FormatBool(qrForcerm)
		if qForcerm != "" {

			if err := r.SetQueryParam("forcerm", qForcerm); err != nil {
				return err
			}
		}
	}

	if o.Httpproxy != nil {

		// query param httpproxy
		var qrHttpproxy bool

		if o.Httpproxy != nil {
			qrHttpproxy = *o.Httpproxy
		}
		qHttpproxy := swag.FormatBool(qrHttpproxy)
		if qHttpproxy != "" {

			if err := r.SetQueryParam("httpproxy", qHttpproxy); err != nil {
				return err
			}
		}
	}

	if o.Labels != nil {

		// query param labels
		var qrLabels string

		if o.Labels != nil {
			qrLabels = *o.Labels
		}
		qLabels := qrLabels
		if qLabels != "" {

			if err := r.SetQueryParam("labels", qLabels); err != nil {
				return err
			}
		}
	}

	if o.Layers != nil {

		// query param layers
		var qrLayers bool

		if o.Layers != nil {
			qrLayers = *o.Layers
		}
		qLayers := swag.FormatBool(qrLayers)
		if qLayers != "" {

			if err := r.SetQueryParam("layers", qLayers); err != nil {
				return err
			}
		}
	}

	if o.Memory != nil {

		// query param memory
		var qrMemory int64

		if o.Memory != nil {
			qrMemory = *o.Memory
		}
		qMemory := swag.FormatInt64(qrMemory)
		if qMemory != "" {

			if err := r.SetQueryParam("memory", qMemory); err != nil {
				return err
			}
		}
	}

	if o.Memswap != nil {

		// query param memswap
		var qrMemswap int64

		if o.Memswap != nil {
			qrMemswap = *o.Memswap
		}
		qMemswap := swag.FormatInt64(qrMemswap)
		if qMemswap != "" {

			if err := r.SetQueryParam("memswap", qMemswap); err != nil {
				return err
			}
		}
	}

	if o.Networkmode != nil {

		// query param networkmode
		var qrNetworkmode string

		if o.Networkmode != nil {
			qrNetworkmode = *o.Networkmode
		}
		qNetworkmode := qrNetworkmode
		if qNetworkmode != "" {

			if err := r.SetQueryParam("networkmode", qNetworkmode); err != nil {
				return err
			}
		}
	}

	if o.Nocache != nil {

		// query param nocache
		var qrNocache bool

		if o.Nocache != nil {
			qrNocache = *o.Nocache
		}
		qNocache := swag.FormatBool(qrNocache)
		if qNocache != "" {

			if err := r.SetQueryParam("nocache", qNocache); err != nil {
				return err
			}
		}
	}

	if o.Outputs != nil {

		// query param outputs
		var qrOutputs string

		if o.Outputs != nil {
			qrOutputs = *o.Outputs
		}
		qOutputs := qrOutputs
		if qOutputs != "" {

			if err := r.SetQueryParam("outputs", qOutputs); err != nil {
				return err
			}
		}
	}

	if o.Platform != nil {

		// query param platform
		var qrPlatform string

		if o.Platform != nil {
			qrPlatform = *o.Platform
		}
		qPlatform := qrPlatform
		if qPlatform != "" {

			if err := r.SetQueryParam("platform", qPlatform); err != nil {
				return err
			}
		}
	}

	if o.Pull != nil {

		// query param pull
		var qrPull bool

		if o.Pull != nil {
			qrPull = *o.Pull
		}
		qPull := swag.FormatBool(qrPull)
		if qPull != "" {

			if err := r.SetQueryParam("pull", qPull); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ bool

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := swag.FormatBool(qrQ)
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Remote != nil {

		// query param remote
		var qrRemote string

		if o.Remote != nil {
			qrRemote = *o.Remote
		}
		qRemote := qrRemote
		if qRemote != "" {

			if err := r.SetQueryParam("remote", qRemote); err != nil {
				return err
			}
		}
	}

	if o.Rm != nil {

		// query param rm
		var qrRm bool

		if o.Rm != nil {
			qrRm = *o.Rm
		}
		qRm := swag.FormatBool(qrRm)
		if qRm != "" {

			if err := r.SetQueryParam("rm", qRm); err != nil {
				return err
			}
		}
	}

	if o.Shmsize != nil {

		// query param shmsize
		var qrShmsize int64

		if o.Shmsize != nil {
			qrShmsize = *o.Shmsize
		}
		qShmsize := swag.FormatInt64(qrShmsize)
		if qShmsize != "" {

			if err := r.SetQueryParam("shmsize", qShmsize); err != nil {
				return err
			}
		}
	}

	if o.Squash != nil {

		// query param squash
		var qrSquash bool

		if o.Squash != nil {
			qrSquash = *o.Squash
		}
		qSquash := swag.FormatBool(qrSquash)
		if qSquash != "" {

			if err := r.SetQueryParam("squash", qSquash); err != nil {
				return err
			}
		}
	}

	if o.T != nil {

		// query param t
		var qrT string

		if o.T != nil {
			qrT = *o.T
		}
		qT := qrT
		if qT != "" {

			if err := r.SetQueryParam("t", qT); err != nil {
				return err
			}
		}
	}

	if o.Target != nil {

		// query param target
		var qrTarget string

		if o.Target != nil {
			qrTarget = *o.Target
		}
		qTarget := qrTarget
		if qTarget != "" {

			if err := r.SetQueryParam("target", qTarget); err != nil {
				return err
			}
		}
	}

	if o.Unsetenv != nil {

		// binding items for unsetenv
		joinedUnsetenv := o.bindParamUnsetenv(reg)

		// query array param unsetenv
		if err := r.SetQueryParam("unsetenv", joinedUnsetenv...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamImageBuildLibpod binds the parameter unsetenv
func (o *ImageBuildLibpodParams) bindParamUnsetenv(formats strfmt.Registry) []string {
	unsetenvIR := o.Unsetenv

	var unsetenvIC []string
	for _, unsetenvIIR := range unsetenvIR { // explode []string

		unsetenvIIV := unsetenvIIR // string as string
		unsetenvIC = append(unsetenvIC, unsetenvIIV)
	}

	// items.CollectionFormat: ""
	unsetenvIS := swag.JoinByFormat(unsetenvIC, "")

	return unsetenvIS
}
