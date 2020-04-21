package us

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           RedirectService
}

func NewInstrumentingMiddleware(requestCount metrics.Counter,
	requestLatency metrics.Histogram,
	next RedirectService) RedirectService {
	return &instrumentingMiddleware{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		next:           next,
	}
}

func (mw instrumentingMiddleware) Find(code string) (url string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	url, err = mw.next.Find(code)
	return
}

func (mw instrumentingMiddleware) Store(url string) (code string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "find", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	code, err = mw.next.Store(url)
	return
}
