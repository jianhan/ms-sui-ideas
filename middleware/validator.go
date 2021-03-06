package middleware

import (
	"context"

	"github.com/jianhan/pkg/http"
	"github.com/micro/go-micro/server"
	"github.com/sirupsen/logrus"
)

// Validator is a handler wrapper for validation of request
func Validator(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		if m, ok := interface{}(req.Request()).(http.Validator); ok {
			logrus.Info("Validation")
			if err := m.Validate(); err != nil {
				return err
			}
		}

		return fn(ctx, req, rsp)
	}
}
