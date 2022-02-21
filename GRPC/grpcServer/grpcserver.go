package grpcServer

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/service"
)

var logger = logging.GetLogger()

type GRPCServer struct {
	service *service.Service
	auth_proto.UnimplementedAuthServer
}

func NewGRPCServer(service *service.Service) {
	s := grpc.NewServer()
	str := &GRPCServer{service: service}
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
func (g *GRPCServer) GetUserWithRights(ctx context.Context, request *auth_proto.AccessToken) (*auth_proto.Response, error) {

	return &auth_proto.Response{
		UserId:      1,
		Role:        1,
		Permissions: request.AccessToken,
	}, nil
}
func (g *GRPCServer) CheckToken(context.Context, *auth_proto.AccessToken) (*auth_proto.Result, error) {
	return &auth_proto.Result{
		Result: true,
	}, nil
}
func (g *GRPCServer) BindUserAndRole(ctx context.Context, user *auth_proto.User) (*auth_proto.Resp, error) {
	return &auth_proto.Resp{}, g.service.BindUserWithRole(user)
}

func (g *GRPCServer) TokenGenerationByRefresh(context.Context, *auth_proto.RefreshToken) (*auth_proto.GeneratedTokens, error) {
	return nil, nil
}
func (g *GRPCServer) TokenGenerationById(ctx context.Context, user *auth_proto.User) (*auth_proto.GeneratedTokens, error) {
	if user.RoleId == 0 {
		role, err := g.service.GetRoleById(int(user.UserId))
		if err != nil {
			logger.Errorf("GetRoleById:%s", err)
			return nil, fmt.Errorf("GetRoleById:%w", err)
		}
		user.RoleId = int32(role.ID)
	} else {
		err := g.service.BindUserWithRole(user)
		if err != nil {
			logger.Errorf("BindUserWithRole:%s", err)
			return nil, fmt.Errorf("BindUserWithRole:%w", err)
		}
	}
	return &auth_proto.GeneratedTokens{
		AccessToken:  fmt.Sprintf("UserId:%d, RoleId:%d", user.UserId, user.RoleId),
		RefreshToken: "Refresh Token",
	}, nil
}

func (g *GRPCServer) GetSalt(context.Context, *auth_proto.ReqSalt) (*auth_proto.Salt, error) {
	return nil, nil
}
