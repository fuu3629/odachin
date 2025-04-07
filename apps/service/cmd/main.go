package main

import (
	"log"
	"net"

	database "github.com/fuu3629/odachin/apps/service/internal/db"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// protoのコンパイルseviceで
// buf generate proto

func main() {
	err := godotenv.Load()
	//dbの準備
	db := database.DbConn()
	//マイグレーションの実行 これカラムの変更とか削除は対応されないため注意。運用段階になった時考えるけど、今のところdbに変更があった場合一回全部消すのが楽そう
	database.Migrations(db)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	presentation.NewServer(grpcServer, db)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
