package us

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingMiddleware struct {
	logger log.Logger
	next   RedirectService
}

func NewLoggingMiddleware(logger log.Logger, next RedirectService) RedirectService {
	return &loggingMiddleware{logger: logger, next: next}
}

func (mw loggingMiddleware) Find(code string) (url string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "find",
			"input", code,
			"output", url,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	url, err = mw.next.Find(code)
	return
}

func (mw loggingMiddleware) Store(url string) (code string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "store",
			"input", url,
			"output", code,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	code, err = mw.next.Store(url)
	return
}
