package main

import (
	"github.com/gadieichhorn/go-kit-url-shortener/pkg/us"
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	logger := log.NewLogfmtLogger(os.Stderr)

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "url_shortener",
		Subsystem: "find_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "url_shortener",
		Subsystem: "find_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	// countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
	// 	Namespace: "my_group",
	// 	Subsystem: "string_service",
	// 	Name:      "count_result",
	// 	Help:      "The result of each count method.",
	// }, []string{}) // no fields here

	var svc us.RedirectService
	svc = us.NewRedirectService(us.NewRedirectRepository())
	svc = us.NewLoggingMiddleware(logger, svc)
	svc = us.NewInstrumentingMiddleware(requestCount, requestLatency, svc)

	findHandler := httptransport.NewServer(
		us.MakeFindEndpoint(svc),
		us.DecodeFindRequest,
		us.EncodeResponse,
	)

	storeHandler := httptransport.NewServer(
		us.MakeStoreEndpoint(svc),
		us.DecodeStoreRequest,
		us.EncodeResponse,
	)

	http.Handle("/store", storeHandler)
	http.Handle("/find", findHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))

}
