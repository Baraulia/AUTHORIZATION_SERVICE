package grpcServer

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/service"
	"strings"
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
	return g.service.AuthUser.ParseToken(request.AccessToken)
}

func (g *GRPCServer) BindUserAndRole(ctx context.Context, user *authProto.User) (*authProto.ResultBinding, error) {
	res, err := g.service.AuthUser.AddRoleToUser(user)
	if err != nil {
		return nil, err
	}
	return &authProto.ResultBinding{Result: res}, nil
}

func (g *GRPCServer) TokenGenerationByRefresh(ctx context.Context, token *authProto.RefreshToken) (*authProto.GeneratedTokens, error) {
	return g.service.RefreshTokens(token.RefreshToken)
}

func (g *GRPCServer) TokenGenerationByUserId(ctx context.Context, user *authProto.User) (*authProto.GeneratedTokens, error) {
	return g.service.GenerateTokensByAuthUser(user)
}

func (g *GRPCServer) GetAllRoles(context.Context, *empty.Empty) (*authProto.Roles, error) {
	roles, err := g.service.GetAllRoles()
	if err != nil {
		return nil, fmt.Errorf("GetAllRoles: %w", err)
	}
	var sliceRoles []string
	for _, role := range roles {
		sliceRoles = append(sliceRoles, role.Name)
	}
	return &authProto.Roles{Roles: strings.Join(sliceRoles[:], ",")}, nil
}
