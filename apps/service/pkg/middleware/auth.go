package middleware

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"connectrpc.com/connect"
	"github.com/dgrijalva/jwt-go"
)

const tokenHeader = "authorization"

var authNotRequiredInterceptorMethods = []string{
	"/odachin.auth.AuthService/CreateUser",
	"/odachin.auth.AuthService/Login",
	"/odachin.allowance.AllowanceService/Allowance",
}

// TODO ROLEによる認可を実装する
func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				return next(ctx, req)
			}
			if slices.Contains(authNotRequiredInterceptorMethods, req.Spec().Procedure) {
				return next(ctx, req)
			}
			authHeader := req.Header().Get(tokenHeader)
			if authHeader == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no authorization header"))
			}
			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid authorization header format"))
			}
			tokenString := parts[1]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(os.Getenv("JWT_SECRET_KEY")), nil
			})
			if err != nil || !token.Valid {
				return nil, connect.NewError(connect.CodeUnauthenticated, fmt.Errorf("invalid token: %w", err))
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid token claims"))
			}

			userID, ok := claims["user_id"]
			if !ok {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("user_id not found in token"))
			}
			ctx = context.WithValue(ctx, "user_id", userID)
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
