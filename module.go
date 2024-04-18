package zeromq

import (
	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type RootModule struct{}
type ZeroMQMetrics struct {
	RequestDuration    *metrics.Metric
	RequestCount       *metrics.Metric
	FailedRequestCount *metrics.Metric
}
type ZeroMQ struct {
	vu      modules.VU
	metrics ZeroMQMetrics
}

func init() {
	modules.Register("k6/x/zeromq", New())
}

func New() *RootModule {
	return &RootModule{}
}

func (*RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &ZeroMQ{
		vu:      vu,
		metrics: registerMetrics(vu),
	}
}

func (z *ZeroMQ) Exports() modules.Exports {
	return modules.Exports{Default: z}
}
