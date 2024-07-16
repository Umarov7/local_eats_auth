package pkg

import (
	"auth-service/config"
	pba "auth-service/genproto/auth"
	"errors"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func CreateAuthClient(cfg *config.Config) pba.AuthClient {
	conn, err := grpc.NewClient(cfg.AUTH_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Println(errors.New("failed to connect to the address: " + err.Error()))
		return nil
	}

	return pba.NewAuthClient(conn)
}
