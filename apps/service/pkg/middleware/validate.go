package middleware

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/proto"
)

func NewValidateInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			v, err := protovalidate.New()
			if err != nil {
				return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to initialize validator: %v", err))
			}
			msg, ok := req.Any().(proto.Message)
			if !ok {
				return nil, errors.New("failed to type assertion proto.Message")
			}
			if err = v.Validate(msg); err != nil {
				return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("validation failed: %v", err))
			}
			return next(ctx, req)

		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
