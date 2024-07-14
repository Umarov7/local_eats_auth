package main

import (
	"auth-service/config"
	pbk "auth-service/genproto/kitchen"
	pbu "auth-service/genproto/user"
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

	server := grpc.NewServer()
	pbu.RegisterUserServer(server, userService)
	pbk.RegisterKitchenServer(server, kitchenService)

	log.Printf("server listening at %v", lis.Addr())
	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("error while serving: %v", err)
	}
}
