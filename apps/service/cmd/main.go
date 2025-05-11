package main

import (
	"log"
	"net/http"
	"os"

	database "github.com/fuu3629/odachin/apps/service/internal/db"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

//TODO 自動デプロイ作る
//TODO 本番環境のDBの接続先を考える。

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	//dbの準備
	db := database.DbConn()
	//マイグレーションの実行 これカラムの変更とか削除は対応されないため注意。運用段階になった時考えるけど、今のところdbに変更があった場合一回全部消すのが楽そう
	database.Migrations(db)

	//開発用のモックデータの投入
	database.Seed(db)

	mux := http.NewServeMux()
	presentation.NewServer(mux, db)

	http.ListenAndServe(
		os.Getenv("PORT"),
		cors.AllowAll().Handler(
			// HTTP1.1リクエストはHTTP/2にアップグレードする
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
}
