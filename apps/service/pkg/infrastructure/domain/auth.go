package domain

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/metadata"
)

func GenerateToken(userId string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // 72時間が有効期限
	}

	// ヘッダーとペイロード生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return accessToken, nil
}

func ValidateToken(tokenString string) (string, error) {
	claims := &jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", err
	}
	return (*claims)["user_id"].(string), nil
}

func ExtractTokenMetadata(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("metadata is not provided")
	}
	tokenString := md.Get("auth")
	if len(tokenString) == 0 {
		return "", fmt.Errorf("no auth token provided")
	}

	token, _ := jwt.Parse(tokenString[0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	if ok && token.Valid {
		fmt.Printf("user_id: %v\n", string(claims["user_id"].(string)))
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
		return claims["user_id"].(string), nil
	} else {
		return "", fmt.Errorf("no userId")
	}
}
