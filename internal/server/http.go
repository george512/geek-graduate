package server

import (
	"context"
	"geek-graduate/configs"
	"geek-graduate/internal/pkg/transport/http"
	"geek-graduate/pb"
	"google.golang.org/grpc"
	"log"
)

func NewHttpServer(
	c *configs.Conf,
) (*http.Server, error) {

	if c.HttpAddress == "" {
		c.HttpAddress = "0.0.0.0:8084"
	}

	if c.EndPoint == "" {
		c.EndPoint = "0.0.0.0:8085"
	}
	httpServer := http.NewServer(c.HttpAddress)
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}

	// in-process handler
	err := pb.RegisterAuthServiceHandlerFromEndpoint(context.Background(), httpServer.Mux, c.EndPoint, dialOptions)
	if err != nil {
		return nil, err
	}

	err = pb.RegisterCreateLaptopHandlerFromEndpoint(context.Background(), httpServer.Mux, c.EndPoint, dialOptions)
	if err != nil {
		return nil, err
	}
	log.Printf("Start Http server at %s, endPoint at %s", c.HttpAddress, c.EndPoint)

	return httpServer, nil
}