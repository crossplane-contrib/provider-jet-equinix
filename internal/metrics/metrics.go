package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"

	k8smetrics "sigs.k8s.io/controller-runtime/pkg/metrics"
)

var (
	metricsEquinixApiResponseCodesTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "equinix_api_responses_total",
		Help: "Number of Equinix API responses by return code and first 5 digits of the token",
	}, []string{"service", "method", "code", "account"})

	metricsEquinixApiResponseCodesLast5m = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "equinix_api_responses_last_5m",
		Help: "Number of Equinix API responses by return code and first 5 digits of the token",
	}, []string{"service", "method", "code", "account"})
)

// SetupMetrics will register the known Prometheus metrics with controller-runtime's metrics registry
func SetupMetrics() error {
	k8smetrics.Registry.MustRegister(
		metricsEquinixApiResponseCodesTotal,
		metricsEquinixApiResponseCodesLast5m,
	)

	go func() {
		// Reset the counters every 5 minutes
		ticker := time.NewTicker(5 * time.Minute)
		for range ticker.C {
			metricsEquinixApiResponseCodesLast5m.Reset()
		}
	}()

	return nil
}

// IncEquinixAPIResp will increment the equinix_api_responses_total metric for the specified service, operation, and responseCode tuple
func IncEquinixAPIResp(service, method, code, account string) error {
	metricsEquinixApiResponseCodesLast5m.WithLabelValues(service, method, code, account).Inc()
	metricsEquinixApiResponseCodesTotal.WithLabelValues(service, method, code, account).Inc()
	return nil
}
