package zeromq

import (
	"time"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

func registerMetrics(vu modules.VU) ZeroMQMetrics {
	registry := vu.InitEnv().Registry
	metrics := ZeroMQMetrics{
		RequestDuration:    registry.MustNewMetric("zeromq_req_duration", metrics.Trend, metrics.Time),
		RequestCount:       registry.MustNewMetric("zeromq_req_count", metrics.Counter, metrics.Default),
		FailedRequestCount: registry.MustNewMetric("zeromq_req_failed", metrics.Rate, metrics.Default),
	}
	return metrics
}

func (z *ZeroMQ) reportMetric(metric *metrics.Metric, now time.Time, value float64) {
	state := z.vu.State()
	ctx := z.vu.Context()
	if state == nil || ctx == nil {
		return
	}

	metrics.PushIfNotDone(ctx, state.Samples, metrics.Sample{
		Time: now,
		TimeSeries: metrics.TimeSeries{
			Metric: metric,
		},
		Value: value,
	})
}
