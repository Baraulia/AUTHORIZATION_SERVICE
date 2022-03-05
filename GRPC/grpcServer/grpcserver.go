package grpcServer

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/service"
)

var logger = logging.GetLogger()

type GRPCServer struct {
	service *service.Service
	authProto.UnimplementedAuthServer
}

func NewGRPCServer(service *service.Service) {
	s := grpc.NewServer()
	str := &GRPCServer{service: service}
	authProto.RegisterAuthServer(s, str)
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		logger.Fatalf("NewGRPCServer, Listen:%s", err)
	}
	reflection.Register(s)
	if err = s.Serve(lis); err != nil {
		logger.Fatalf("NewGRPCServer, Serve:%s", err)
	}

}
func (g *GRPCServer) GetUserWithRights(ctx context.Context, request *authProto.AccessToken) (*authProto.UserRole, error) {
	return g.service.Authorization.ParseToken(request.AccessToken)
}

func (g *GRPCServer) BindUserAndRole(ctx context.Context, user *authProto.User) (*authProto.ResultBinding, error) {
	res, err := g.service.RolePerm.AddRoleToUser(user)
	if err != nil {
		return nil, err
	}
	return &authProto.ResultBinding{Result: res}, nil
}

func (g *GRPCServer) TokenGenerationByRefresh(ctx context.Context, token *authProto.RefreshToken) (*authProto.GeneratedTokens, error) {
	return g.service.RefreshTokens(token.RefreshToken)
}
func (g *GRPCServer) TokenGenerationByUserId(ctx context.Context, user *authProto.User) (*authProto.GeneratedTokens, error) {
	if user.RoleId == 0 {
		roleId, err := g.service.GetRoleByUserId(int(user.UserId))
		if err != nil {
			logger.Errorf("GetRoleById:%s", err)
			return nil, fmt.Errorf("GetRoleById:%w", err)
		}
		user.RoleId = int32(roleId)
	} else {
		_, err := g.service.AddRoleToUser(user)
		if err != nil {
			logger.Errorf("BindUserWithRole:%s", err)
			return nil, fmt.Errorf("BindUserWithRole:%w", err)
		}
	}
	return g.service.GenerateTokensByAuthUser(user)
}

func (g *GRPCServer) GetSalt(context.Context, *authProto.ReqSalt) (*authProto.Salt, error) {
	return nil, nil
}
