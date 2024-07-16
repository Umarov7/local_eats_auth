package handler

import (
	pb "auth-service/genproto/auth"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Register godoc
// @Summary Registers user
// @Description Inserts new user into database
// @Tags auth
// @Param data body auth.RegisterRequest true "New user data"
// @Success 200 {object} auth.RegisterResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /auth/register [post]
func (h Handler) Register(c *gin.Context) {
	h.Log.Info("Register function is starting")

	var data pb.RegisterRequest
	if err := c.ShouldBind(&data); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	user, err := h.Auth.Register(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error registering user").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Register has successfully finished")
	c.JSON(http.StatusCreated, gin.H{"New user": user})
}

// Login godoc
// @Summary Logs user in
// @Description Logs user in
// @Tags auth
// @Param data body auth.LoginRequest true "User credentials"
// @Success 200 {object} auth.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /auth/login [post]
func (h Handler) Login(c *gin.Context) {
	h.Log.Info("Login function is starting")

	var data pb.LoginRequest
	if err := c.ShouldBind(&data); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	tokens, err := h.Auth.Login(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error logging in user").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Login has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Tokens": tokens})
}

// Logout godoc
// @Summary Logs user out
// @Description Logs user out
// @Tags auth
// @Param data body auth.Token true "Refresh token"
// @Success 200 {object} string
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /auth/logout [post]
func (h Handler) Logout(c *gin.Context) {
	h.Log.Info("Logout function is starting")

	var data pb.Token
	if err := c.ShouldBind(&data); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	expiredToken, err := h.Auth.Logout(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error logging out user").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Logout has successfully finished")
	c.JSON(http.StatusOK, gin.H{"User logged out successfully": expiredToken})
}

// Refresh godoc
// @Summary Refreshes refresh token
// @Description Refreshes refresh token
// @Tags auth
// @Param data body auth.Token true "Refresh token"
// @Success 200 {object} auth.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /auth/refresh-token [post]
func (h Handler) Refresh(c *gin.Context) {
	h.Log.Info("Refresh function is starting")

	var data pb.Token
	if err := c.ShouldBind(&data); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	tokens, err := h.Auth.RefreshToken(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error refreshing token").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Refresh has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Tokens": tokens})
}

// ForgotPassword godoc
// @Summary Sends reset code
// @Description Sends reset code to user's email
// @Tags auth
// @Param data body auth.ResetRequest true "Email"
// @Success 200 {object} auth.ResetResponse
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /auth/forgot-password [post]
func (h Handler) ForgotPassword(c *gin.Context) {
	h.Log.Info("ForgotPassword function is starting")

	var data pb.ResetRequest
	if err := c.ShouldBind(&data); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	msg, err := h.Auth.ForgotPassword(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error resetting password").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("ForgotPassword has successfully finished")
	c.JSON(http.StatusOK, msg.Message)
}

// ResetPassword godoc
// @Summary Resets password
// @Description Resets password based on reset code
// @Tags auth
// @Param data body auth.Code true "Details"
// @Success 200 {object} auth.Status
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /auth/reset-password [post]
func (h Handler) ResetPassword(c *gin.Context) {
	h.Log.Info("ResetPassword function is starting")

	var data pb.Code
	if err := c.ShouldBind(&data); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	status, err := h.Auth.ResetPassword(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error resetting password").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("ResetPassword has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Password reset": status.Successful})
}
