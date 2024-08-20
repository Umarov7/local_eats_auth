package main

import (
	"auth-service/api"
	"auth-service/config"
	pba "auth-service/genproto/auth"
	pbk "auth-service/genproto/kitchen"
	pbu "auth-service/genproto/user"
	"auth-service/pkg"
	"auth-service/service"
	"auth-service/storage/postgres"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	db, err := postgres.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", cfg.AUTH_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	userService := service.NewUserService(db)
	kitchenService := service.NewKitchenService(db)
	authService := service.NewAuthService(db, cfg)
	authClient := pkg.CreateAuthClient(cfg)

	router := api.Router(authClient)
	log.Printf("Auth api is running on port %v", cfg.HTTP_PORT)
	go router.Run(cfg.HTTP_PORT)

	server := grpc.NewServer()
	pbu.RegisterUserServer(server, userService)
	pbk.RegisterKitchenServer(server, kitchenService)
	pba.RegisterAuthServer(server, authService)

	log.Printf("Auth server is listening at %v", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
