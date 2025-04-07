package presentation

import (
	"github.com/fuu3629/odachin/apps/service/gen/v1/odachin"
	"github.com/fuu3629/odachin/apps/service/pkg/usecase"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type ServerStruct struct {
	useCase usecase.UseCaseImpl
	odachin.UnimplementedOdachinServiceServer
}

func NewServer(grpcServer *grpc.Server, db *gorm.DB) {
	userGrpc := &ServerStruct{useCase: usecase.New(db)}
	odachin.RegisterOdachinServiceServer(grpcServer, userGrpc)
}

// func (s *ServerStruct) GetUser(ctx context.Context, req *odachin.CreateUserRequest) (*odachin.CreateUserResponse, error) {
// 	// Implement the logic to get user details using the use case
// 	// For example:
// 	// user, err := s.useCase.GetUser(req.Id)
// 	// if err != nil {
// 	//     return nil, err
// 	// }
// 	// return &odachin.GetUserResponse{User: user}, nil

// 	return nil, nil // Placeholder return
// }
