package main

import (
	"flag"
	"fmt"
	"geek-graduate/configs"
	"geek-graduate/internal/data"
	"geek-graduate/internal/pkg/app"
	newgrpc "geek-graduate/internal/pkg/transport/grpc"
	newhttp "geek-graduate/internal/pkg/transport/http"
	"geek-graduate/internal/server"
	"geek-graduate/internal/service"
	"log"
	"time"
)

const (
	secretKey     = "secret"
	tokenDuration = 30 * time.Minute
)

func seedUser(userStore data.UserStore) error {
	err := createUser(userStore, "admin1", "secret", "admin")
	if err != nil {
		return err
	}
	return createUser(userStore, "user1", "secret", "user")
}

func createUser(userStore data.UserStore, username, password, role string) error {
	user, err := data.NewUser(username, password, role)
	if err != nil {
		return err
	}
	return userStore.Save(user)
}

func main() {
	flag.Parse()

	c := configs.NewConf()
	if err := c.Load(); err != nil {
		panic(err)
	}

	userStore := data.NewInMemoryUserStore()
	err := seedUser(userStore)
	if err != nil {
		log.Fatal("cannot seed users")
	}
	jwtManger := service.NewJWTManager(secretKey, tokenDuration)
	authServer := service.NewAuthServer(userStore, jwtManger)

	// 创建LaptopServer和grpcServer实例并注册
	laptopServer := service.NewLaptopServer(data.NewInMemoryLaptopStore(), data.NewDiskImageStore("img/"), data.NewInMemoryRatingStore())

	grpcServer := server.NewGRPCServer(c, authServer, laptopServer, jwtManger)
	httpServer, err := server.NewHttpServer(c)
	if err != nil {
		panic(err)
	}

	appli := newApp(httpServer, grpcServer)

	if err := appli.Run(); err != nil {
		fmt.Println(err)
	}
}

func newApp(hs *newhttp.Server, gs *newgrpc.Server) *app.App {
	return app.New(
		hs,
		gs,
	)
}
