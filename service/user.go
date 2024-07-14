package service

import (
	pb "auth-service/genproto/user"
	"auth-service/pkg/logger"
	"auth-service/storage/postgres"
	"context"
	"database/sql"
	"log/slog"

	"github.com/pkg/errors"
)

type UserService struct {
	pb.UnimplementedUserServer
	Repo   *postgres.UserRepo
	Logger *slog.Logger
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		Repo:   postgres.NewUserRepo(db),
		Logger: logger.NewLogger(),
	}
}

func (u *UserService) GetProfile(ctx context.Context, req *pb.ID) (*pb.Profile, error) {
	resp, err := u.Repo.Read(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get profile")
	}

	return resp, nil
}

func (u *UserService) UpdateProfile(ctx context.Context, req *pb.NewInfo) (*pb.Details, error) {
	resp, err := u.Repo.Update(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update profile")
	}

	return resp, nil
}

func (u *UserService) DeleteProfile(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := u.Repo.Delete(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete profile")
	}

	return &pb.Void{}, nil
}
