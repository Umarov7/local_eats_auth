package service

import (
	pb "auth-service/genproto/kitchen"
	"auth-service/pkg/logger"
	"auth-service/storage/postgres"
	"context"
	"database/sql"
	"log/slog"

	"github.com/pkg/errors"
)

type KitchenService struct {
	pb.UnimplementedKitchenServer
	Repo   *postgres.KitchenRepo
	Logger *slog.Logger
}

func NewKitchenService(db *sql.DB) *KitchenService {
	return &KitchenService{
		Repo:   postgres.NewKitchenRepo(db),
		Logger: logger.NewLogger(),
	}
}

func (k *KitchenService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	resp, err := k.Repo.Create(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create kitchen")
	}

	return resp, nil
}

func (k *KitchenService) Get(ctx context.Context, req *pb.ID) (*pb.Info, error) {
	resp, err := k.Repo.Read(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get kitchen")
	}

	return resp, nil
}

func (k *KitchenService) Update(ctx context.Context, req *pb.NewData) (*pb.UpdatedData, error) {
	resp, err := k.Repo.Update(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update kitchen")
	}

	return resp, nil
}

func (k *KitchenService) Delete(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	err := k.Repo.Delete(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete kitchen")
	}

	return &pb.Void{}, nil
}

func (k *KitchenService) Fetch(ctx context.Context, req *pb.Pagination) (*pb.Kitchens, error) {
	resp, total, err := k.Repo.Fetch(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch kitchens")
	}

	return &pb.Kitchens{
		Kitchens: resp,
		Total:    int32(total),
		Page:     req.Offset / req.Limit,
		Limit:    req.Limit,
	}, nil
}

func (k *KitchenService) Search(ctx context.Context, req *pb.SearchDetails) (*pb.Kitchens, error) {
	resp, total, err := k.Repo.Search(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find kitchen")
	}

	return &pb.Kitchens{
		Kitchens: resp,
		Total:    int32(total),
		Page:     req.Pagination.Offset / req.Pagination.Limit,
		Limit:    req.Pagination.Limit,
	}, nil
}
