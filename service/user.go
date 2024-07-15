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
	u.Logger.Info("GetProfile method is starting")

	resp, err := u.Repo.Read(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to get profile")
		u.Logger.Error(er.Error())
		return nil, er
	}

	u.Logger.Info("GetProfile has successfully finished")
	return resp, nil
}

func (u *UserService) UpdateProfile(ctx context.Context, req *pb.NewInfo) (*pb.Details, error) {
	u.Logger.Info("UpdateProfile method is starting")

	resp, err := u.Repo.Update(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to update profile")
		u.Logger.Error(er.Error())
		return nil, er
	}

	u.Logger.Info("UpdateProfile has successfully finished")
	return resp, nil
}

func (u *UserService) DeleteProfile(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	u.Logger.Info("DeleteProfile method is starting")

	err := u.Repo.Delete(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to delete profile")
		u.Logger.Error(er.Error())
		return nil, er
	}

	u.Logger.Info("DeleteProfile has successfully finished")
	return &pb.Void{}, nil
}

func (u *UserService) ValidateUser(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	u.Logger.Info("ValidateUser method is starting")

	exists, err := u.Repo.ValidateUser(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to validate user")
		u.Logger.Error(er.Error())
		return nil, er
	}

	u.Logger.Info("ValidateUser has successfully finished")
	return &pb.Status{Exists: exists}, nil
}
