package middleware

import (
	"context"
	"fmt"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "auth")
	if err != nil {
		return nil, err
	}
	fmt.Printf("receive token: %s\n", token)
	if token != "hoge" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid token")
	}
	newCtx := context.WithValue(ctx, "result", "ok")
	return newCtx, nil
}
