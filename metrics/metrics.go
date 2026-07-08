package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "conduit",
			Name:      "http_requests_total",
			Help:      "Total number of HTTP requests.",
		},
		[]string{"method", "route", "status"},
	)

	HttpRequestsDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "conduit",
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request Latency.",
			Buckets:   prometheus.DefBuckets,
		},
		[]string{"method", "route"},
	)

	InflightRequests = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "conduit",
			Name:      "http_requests_in_flight",
			Help:      "Current requests being processed.",
		},
	)
)
