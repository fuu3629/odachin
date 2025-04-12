package middleware

import (
	"context"
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
)

func AuthFunc(ctx context.Context) (context.Context, error) {
	fmt.Println("AuthFunc")
	tokenString, err := auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["user_id"])
	if ok && token.Valid {
		newCtx := context.WithValue(ctx, "user_id", claims["user_id"])
		return newCtx, nil
	} else {
		return nil, fmt.Errorf("no userId")
	}
}
