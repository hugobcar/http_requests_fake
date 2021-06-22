package prometheus

import (
	"fmt"
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

type FakeMetrics struct {
	Name  string
	Total prometheus.Gauge
}

var (
	mu              *sync.Mutex
	RegistredGauges map[string]FakeMetrics
)

func init() {
	mu = new(sync.Mutex)
	RegistredGauges = make(map[string]FakeMetrics)
}

// CreateGauges - Create Gauges in Prometheus
func CreateGauges() {
	var gName = "http_server_requests_seconds_count"

	defaultLabels := prometheus.Labels{}

	RegistredGauges["http_server_requests_seconds_count"] = FakeMetrics{
		Name:  gName,
		Total: RegisterGauge(gName, fmt.Sprintf("%s", gName), defaultLabels, ""),
	}
}

// RegisterGauge -- register new prometheus guage metrics
func RegisterGauge(name, metric string, labels prometheus.Labels, help string) prometheus.Gauge {
	if g, found := RegistredGauges[name]; found {
		fmt.Println(g)
	}

	g := prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name:        metric,
			Help:        help,
			ConstLabels: labels,
		})

	if err := prometheus.Register(g); err != nil {
		fmt.Println(err)
	}
	return g
}
