package common

import (
	"bytes"
	"os"
	"sort"
	"time"

	"github.com/skupperproject/skupper/test/utils/base"
	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
)

type AdaptorType string

const AdaptorTCP AdaptorType = "tcp"
const AdaptorHTTP AdaptorType = "http"

type AppSettings map[string]string

func (a AppSettings) AddEnvVar(name string) {
	a[name] = os.Getenv(name)
}

type PerformanceApp struct {
	Name           string        `json:"name"`
	Description    string        `json:"description,omitempty"`
	Service        ServiceInfo   `json:"service"`
	Server         *ServerInfo   `json:"server"`
	Client         *ClientInfo   `json:"client"`
	ThroughputUnit string        `json:"throughputUnit,omitempty"`
	LatencyUnit    time.Duration `json:"latencyUnit,omitempty"`
}

type ServerInfo struct {
	Name       string             `json:"name"`
	Resources  ResourceSettings   `json:"resources,omitempty"`
	Settings   AppSettings        `json:"settings,omitempty"`
	Deployment *appsv1.Deployment `json:"deployment"`
}

type ClientInfo struct {
	Name      string           `json:"name"`
	Resources ResourceSettings `json:"resources,omitempty"`
	Settings  AppSettings      `json:"settings,omitempty"`
	Timeout   time.Duration    `json:"timeout,omitempty"`
	Jobs      []JobInfo        `json:"jobs"`
}

func (c *ClientInfo) JobNames() []string {
	jobNames := []string{}
	for _, job := range c.Jobs {
		jobNames = append(jobNames, job.Name)
	}
	sort.Strings(jobNames)
	return jobNames
}

type JobInfo struct {
	Name    string       `json:"name"`
	Clients int          `json:"clients"`
	Job     *batchv1.Job `json:"job"`
}

type ServiceInfo struct {
	Address  string      `json:"address"`
	Protocol string      `json:"protocol"`
	Adaptor  AdaptorType `json:"adaptor"`
	Port     int         `json:"port"`
}

type PerformanceTest interface {
	App() PerformanceApp
	Validate(serverCluster, clientCluster *base.ClusterContext, job JobInfo) Result
}

type Result struct {
	App        PerformanceApp  `json:"app"`
	Sites      int             `json:"sites"`
	Skupper    SkupperSettings `json:"skupper"`
	Failed     bool            `json:"failed,omitempty"`
	Error      error           `json:"error,omitempty"`
	Job        JobInfo         `json:"job"`
	Throughput float64         `json:"throughput,omitempty"`
	LatencyAvg float64         `json:"latencyAvg,omitempty"`
	Latency50  float64         `json:"latency50,omitempty"`
	Latency99  float64         `json:"latency99,omitempty"`
}

func (r *Result) SetError(err error) {
	if err != nil {
		r.Failed = true
	}
	r.Error = err
}

type SkupperSettings struct {
	Sites   []int          `json:"sites"`
	Timeout string         `json:"timeout,omitempty"`
	Router  RouterSettings `json:"router,omitempty"`
}

type RouterSettings struct {
	MaxFrameSize     int              `json:"maxFrameSize,omitempty"`
	MaxSessionFrames int              `json:"maxSessionFrames,omitempty"`
	Resources        ResourceSettings `json:"resources"`
}

type ResourceSettings struct {
	Memory string `json:"memory,omitempty"`
	CPU    string `json:"CPU,omitempty"`
}

// FlushWriter Buffers content till flushed
type FlushWriter struct {
	buf bytes.Buffer
}

func (f *FlushWriter) Write(p []byte) (n int, err error) {
	return f.buf.Write(p)
}

func (f *FlushWriter) Flush() (n int, err error) {
	return os.Stdout.WriteString(f.buf.String())
}
