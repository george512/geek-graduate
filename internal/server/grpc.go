package server

import (
	"geek-graduate/configs"
	"geek-graduate/internal/service"
	"geek-graduate/pb"

	newgrpc "geek-graduate/internal/pkg/transport/grpc"
	"google.golang.org/grpc"
	"log"
)

func accessibleRoles() map[string][]string {
	const laptopServicePath = "/george.pcbook.CreateLaptop/"
	return map[string][]string{
		laptopServicePath + "CreateLaptop": {"admin"},
		laptopServicePath + "UploadImage":  {"admin"},
		laptopServicePath + "RateLaptop":   {"user"},
		laptopServicePath + "SearchLaptop": {"admin"},
	}
}

func NewGRPCServer(
	c *configs.Conf,
	authServer pb.AuthServiceServer,
	laptopServer pb.CreateLaptopServer,
	jwtManager *service.JWTManager,
) *newgrpc.Server {
	if c.GrpcAddress == "" {
		c.GrpcAddress = "0.0.0.0:8085"
	}
	interceptor := service.NewAuthInterceptor(jwtManager, accessibleRoles())
	serverOption := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	}

	grpcServer := newgrpc.NewServer(c.GrpcAddress, serverOption...)

	pb.RegisterCreateLaptopServer(grpcServer, laptopServer)
	pb.RegisterAuthServiceServer(grpcServer, authServer)

	log.Printf("Start GRPC server at %s ", c.GrpcAddress)
	return grpcServer
}
