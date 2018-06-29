package middleware

import (
	"context"

	"github.com/jianhan/ms-sui-ideas/handler"
	"github.com/micro/go-micro/server"
)

// Validator is a handler wrapper for validation of request
func Validator(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		if m, ok := interface{}(req).(handler.Validator); ok {
			if err := m.Validate(); err != nil {
				return err
			}
		}
		return fn(ctx, req, rsp)
	}
}
