package middleware

import (
	"context"
	"log"

	"github.com/micro/go-micro/server"
)

// LogWrapper is a handler wrapper
func LogWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, rsp interface{}) error {
		log.Printf("[wrapper] server request: %v", req.Method())
		return fn(ctx, req, rsp)
	}
}
