package middleware

import (
	"context"

	"github.com/jianhan/pkg/http"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
)

// Preprocessor is a handler wrapper for preprocess request
func Preprocessor(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		if m, ok := interface{}(req.Request()).(http.Preprocessor); ok {
			logrus.Info("Pre process")
			if err := m.Preprocess(); err != nil {
				return err
			}
		}

		return fn(ctx, req, rsp)
	}
}
