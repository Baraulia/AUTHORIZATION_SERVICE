package grpcServer

import (
	"context"
	"fmt"
	auth_proto "github.com/Baraulia/AUTHORIZATION_SERVICE/GRPC"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/pkg/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

var logger = logging.GetLogger()

type GRPCServer struct{}

func NewGRPCServer() {
	s := grpc.NewServer()
	str := &GRPCServer{}
	auth_proto.RegisterAuthServer(s, str)
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		logger.Fatalf("NewGRPCServer, Listen:%s", err)
	}
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		logger.Fatalf("NewGRPCServer, Serve:%s", err)
	}

}
func (g *GRPCServer) GetUserWithRights(ctx context.Context, request *auth_proto.Request) (*auth_proto.Response, error) {

	return &auth_proto.Response{
		UserId:      1,
		Role:        "Super",
		Permissions: request.AccessToken,
	}, nil
}
func (g *GRPCServer) CheckToken(context.Context, *auth_proto.AccessToken) (*auth_proto.Result, error) {
	return nil, nil
}
func (g *GRPCServer) TokenGenerationByRefresh(context.Context, *auth_proto.RefreshToken) (*auth_proto.GeneratedTokens, error) {
	return nil, nil
}
func (g *GRPCServer) TokenGenerationById(ctx context.Context, user *auth_proto.User) (*auth_proto.GeneratedTokens, error) {
	if user.Role == "" {
		user.Role = "Superadmin"
	}
	return &auth_proto.GeneratedTokens{
		AccessToken:  fmt.Sprintf("UserId:%d", user.UserId),
		RefreshToken: user.Role,
	}, nil
}
func (g *GRPCServer) GetSalt(context.Context, *auth_proto.ReqSalt) (*auth_proto.Salt, error) {
	return nil, nil
}
func (g *GRPCServer) MustEmbedUnimplementedAuthServer() {

}
