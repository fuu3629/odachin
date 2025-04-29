package middleware

import (
	"context"
	"log"
	"time"

	"connectrpc.com/connect"
)

func NewLoggerInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			start := time.Now()
			method := req.Spec().Procedure
			log.Printf("[Request] %s", method)

			for key, vals := range req.Header() {
				log.Printf("Header[%s] = %v", key, vals)
			}

			resp, err := next(ctx, req)

			duration := time.Since(start)
			if err != nil {
				log.Printf("[Error] %s failed: %v (%v)", method, err, duration)
			} else {
				log.Printf("[Success] %s completed in %v", method, duration)
			}
			return resp, err
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
