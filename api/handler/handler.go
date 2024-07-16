package handler

import (
	"auth-service/genproto/auth"
	"auth-service/pkg/logger"
	"log/slog"
)

type Handler struct {
	Auth auth.AuthClient
	Log  *slog.Logger
}

func NewHandler(auth auth.AuthClient) *Handler {
	return &Handler{
		Auth: auth,
		Log:  logger.NewLogger(),
	}
}
