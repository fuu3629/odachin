package main

import (
	"log"
	"net"

	"github.com/bufbuild/protovalidate-go"
	database "github.com/fuu3629/odachin/apps/service/internal/db"
	"github.com/fuu3629/odachin/apps/service/pkg/middleware"
	"github.com/fuu3629/odachin/apps/service/pkg/presentation"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	protovalidate_middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/protovalidate"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//TODO 自動デプロイ作る
//TODO test作る
//TODO transaction系ちゃんとする
//TODO wallet系実装する
//TODO ROLE系の認証

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

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	validator, err := protovalidate.New()
	if err != nil {
		log.Fatalf("failed to create validator: %v", err)
	}
	recovery_opts := []recovery.Option{
		recovery.WithRecoveryHandler(middleware.RecoveryFunc),
	}
	logger := zap.NewExample()

	logger_opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			auth.UnaryServerInterceptor(middleware.AuthFunc),
			protovalidate_middleware.UnaryServerInterceptor(validator),
			recovery.UnaryServerInterceptor(recovery_opts...),
			logging.UnaryServerInterceptor(middleware.InterceptorLogger(logger), logger_opts...),
		),
	)
	presentation.NewServer(grpcServer, db)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
