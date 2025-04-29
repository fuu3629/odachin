package middleware

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func RecoveryFunc(p interface{}) error {
	fmt.Printf("p: %+v\n", p)
	return status.Errorf(codes.Internal, "Unexpected error")
}

func NewRecoveryInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (resp connect.AnyResponse, err error) {
			defer func() {
				if r := recover(); r != nil {
					err = connect.NewError(
						connect.CodeInternal,
						fmt.Errorf("internal server panic: %v", r),
					)
				}
			}()
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
