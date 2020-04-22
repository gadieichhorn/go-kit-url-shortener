package main

import (
	"github.com/gadieichhorn/go-kit-url-shortener/pkg/shortener"
	"net/http"
	"os"

	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

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

	var svc shortener.RedirectService
	svc = shortener.NewRedirectService(shortener.NewRedirectRepository())
	svc = shortener.NewLoggingMiddleware(logger, svc)
	svc = shortener.NewInstrumentingMiddleware(requestCount, requestLatency, svc)

	findHandler := httptransport.NewServer(
		shortener.MakeFindEndpoint(svc),
		shortener.DecodeFindRequest,
		shortener.EncodeResponse,
	)

	storeHandler := httptransport.NewServer(
		shortener.MakeStoreEndpoint(svc),
		shortener.DecodeStoreRequest,
		shortener.EncodeResponse,
	)

	http.Handle("/store", storeHandler)
	http.Handle("/find", findHandler)
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))

}
