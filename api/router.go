package api

import (
	_ "auth-service/api/docs"
	"auth-service/api/handler"
	"auth-service/genproto/auth"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Authorazation
// @version 1.0
// @description Authorazation API
// @host localhost:8081
// BasePath: /
func Router(client auth.AuthClient) *gin.Engine {
	h := handler.NewHandler(client)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := r.Group("/auth")
	auth.POST("/register", h.Register)
	auth.POST("/login", h.Login)
	auth.POST("/forgot-password", h.ForgotPassword)
	auth.POST("/reset-password", h.ResetPassword)
	auth.POST("/refresh-token", h.Refresh)
	auth.POST("/logout", h.Logout)

	return r
}
