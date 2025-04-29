package main

import (
	"log"
	"net/http"

	database "github.com/fuu3629/odachin/apps/service/internal/db"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

//TODO 自動デプロイ作る
//TODO transaction系ちゃんとする
//TODO wallet系実装する

// TODO connect対応にする
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
		"localhost:50051",
		cors.AllowAll().Handler(
			// HTTP1.1リクエストはHTTP/2にアップグレードされる
			h2c.NewHandler(mux, &http2.Server{}),
		),
	)
}
