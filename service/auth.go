package service

import (
	"auth-service/api/token"
	"auth-service/config"
	pb "auth-service/genproto/auth"
	"auth-service/pkg/logger"
	"auth-service/storage/postgres"
	"auth-service/storage/redis"
	"context"
	"database/sql"
	"log/slog"
	"math/rand"
	"strconv"

	"github.com/pkg/errors"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	Repo   *postgres.UserRepo
	Config *config.Config
	Logger *slog.Logger
}

func NewAuthService(db *sql.DB, cfg *config.Config) *AuthService {
	return &AuthService{
		Repo:   postgres.NewUserRepo(db),
		Config: cfg,
		Logger: logger.NewLogger(),
	}
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	s.Logger.Info("Register method is starting")

	resp, err := s.Repo.Create(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to register")
		s.Logger.Error(er.Error())
		return nil, er
	}

	s.Logger.Info("Register has successfully finished")
	return resp, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.Tokens, error) {
	s.Logger.Info("Login method is starting")

	id, username, passwordHash, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		er := errors.Wrap(err, "failed to find user")
		s.Logger.Error(er.Error())
		return nil, er
	}

	if passwordHash != req.Password {
		er := errors.New("incorrect password")
		s.Logger.Error(er.Error())
		return nil, er
	}

	accessToken, err := token.GenerateAccessToken(id, username, req.Email)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	refreshToken, err := token.GenerateRefreshToken(id)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	s.Logger.Info("Login has successfully finished")
	return &pb.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, t *pb.Token) (*pb.Tokens, error) {
	s.Logger.Info("CheckRefreshToken method is starting")

	_, err := token.ValidateRefreshToken(t.RefreshToken)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	id, err := token.GetUserIdFromRefreshToken(t.RefreshToken)
	if err != nil {
		er := errors.Wrap(err, "failed to get user id")
		s.Logger.Error(er.Error())
		return nil, er
	}

	username, email, _, err := s.Repo.GetUserByID(ctx, id)
	if err != nil {
		er := errors.Wrap(err, "failed to get user info")
		s.Logger.Error(er.Error())
		return nil, er
	}

	accessToken, err := token.GenerateAccessToken(id, username, email)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	s.Logger.Info("CheckRefreshToken has successfully finished")
	return &pb.Tokens{
		AccessToken:  accessToken,
		RefreshToken: t.RefreshToken,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, t *pb.Token) (*pb.Token, error) {
	s.Logger.Info("Logout method is starting")

	_, err := token.ValidateRefreshToken(t.RefreshToken)
	if err != nil {
		s.Logger.Error(err.Error())
		return nil, err
	}

	expiredToken, err := token.InvalidateRefreshToken(t.RefreshToken)
	if err != nil {
		er := errors.Wrap(err, "failed to invalidate refresh token")
		s.Logger.Error(er.Error())
		return nil, er
	}

	s.Logger.Info("Logout has successfully finished")
	return &pb.Token{RefreshToken: expiredToken}, nil
}

func (s *AuthService) ForgotPassword(ctx context.Context, req *pb.ResetRequest) (*pb.ResetResponse, error) {
	s.Logger.Info("ForgotPassword method is starting")

	id, _, _, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		er := errors.Wrap(err, "failed to find user")
		s.Logger.Error(er.Error())
		return nil, er
	}

	code := strconv.Itoa(rand.Intn(900000) + 100000)

	err = redis.StoreCode(ctx, id, code)
	if err != nil {
		er := errors.Wrap(err, "failed to store code")
		s.Logger.Error(er.Error())
		return nil, er
	}

	err = SendCode(s.Config, req.Email, code)
	if err != nil {
		er := errors.Wrap(err, "failed to send code")
		s.Logger.Error(er.Error())
		return nil, er
	}

	s.Logger.Info("ForgotPassword has successfully finished")
	return &pb.ResetResponse{Message: "Password reset link has been sent to your email"}, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, req *pb.Code) (*pb.Status, error) {
	s.Logger.Info("ResetPassword method is starting")

	id, _, _, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		er := errors.Wrap(err, "failed to find user")
		s.Logger.Error(er.Error())
		return nil, er
	}

	code, err := redis.GetCode(ctx, id)
	if err != nil {
		er := errors.Wrap(err, "failed to get code")
		s.Logger.Error(er.Error())
		return nil, er
	}

	if code != req.Code {
		er := errors.New("incorrect code")
		s.Logger.Error(er.Error())
		return nil, er
	}

	err = redis.DeleteCode(ctx, id)
	if err != nil {
		er := errors.Wrap(err, "failed to delete code")
		s.Logger.Error(er.Error())
		return nil, er
	}

	err = s.Repo.UpdatePassword(ctx, id, req.Password)
	if err != nil {
		er := errors.Wrap(err, "failed to update password")
		s.Logger.Error(er.Error())
		return nil, er
	}

	s.Logger.Info("ResetPassword has successfully finished")
	return &pb.Status{Successful: true}, nil
}
