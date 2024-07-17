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
	k.Logger.Info("CreateKitchen method is starting")

	resp, err := k.Repo.Create(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to create kitchen")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("CreateKitchen has successfully finished")
	return resp, nil
}

func (k *KitchenService) Get(ctx context.Context, req *pb.ID) (*pb.Info, error) {
	k.Logger.Info("GetKitchen method is starting")

	resp, err := k.Repo.Read(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to get kitchen")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("GetKitchen has successfully finished")
	return resp, nil
}

func (k *KitchenService) Update(ctx context.Context, req *pb.NewData) (*pb.UpdatedData, error) {
	k.Logger.Info("UpdateKitchen method is starting")

	resp, err := k.Repo.Update(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to update kitchen")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("UpdateKitchen has successfully finished")
	return resp, nil
}

func (k *KitchenService) Delete(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	k.Logger.Info("DeleteKitchen method is starting")

	err := k.Repo.Delete(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to delete kitchen")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("DeleteKitchen has successfully finished")
	return &pb.Void{}, nil
}

func (k *KitchenService) Fetch(ctx context.Context, req *pb.Pagination) (*pb.Kitchens, error) {
	k.Logger.Info("FetchKitchens method is starting")

	resp, total, err := k.Repo.Fetch(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to fetch kitchens")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("FetchKitchens has successfully finished")
	return &pb.Kitchens{
		Kitchens: resp,
		Total:    int32(total),
		Page:     req.Offset / req.Limit,
		Limit:    req.Limit,
	}, nil
}

func (k *KitchenService) Search(ctx context.Context, req *pb.SearchDetails) (*pb.Kitchens, error) {
	k.Logger.Info("SearchKitchen method is starting")

	resp, total, err := k.Repo.Search(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "failed to find kitchen")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("SearchKitchen has successfully finished")
	return &pb.Kitchens{
		Kitchens: resp,
		Total:    int32(total),
		Page:     req.Pagination.Offset / req.Pagination.Limit,
		Limit:    req.Pagination.Limit,
	}, nil
}

func (k *KitchenService) ValidateKitchen(ctx context.Context, req *pb.ID) (*pb.Status, error) {
	k.Logger.Info("ValidateKitchen method is starting")

	exists, err := k.Repo.ValidateKitchen(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to validate kitchen")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("ValidateKitchen has successfully finished")
	return &pb.Status{Exists: exists}, nil
}

func (k *KitchenService) IncrementTotalOrders(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	k.Logger.Info("IncrementTotalOrders method is starting")

	err := k.Repo.IncrementTotalOrders(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to increment total orders")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("IncrementTotalOrders has successfully finished")
	return &pb.Void{}, nil
}

func (k *KitchenService) UpdateRating(ctx context.Context, req *pb.Rating) (*pb.Void, error) {
	k.Logger.Info("UpdateRating method is starting")

	err := k.Repo.UpdateRating(ctx, req.Id, req.Rating)
	if err != nil {
		er := errors.Wrap(err, "failed to update rating")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("UpdateRating has successfully finished")
	return &pb.Void{}, nil
}

func (k *KitchenService) UpdateRevenue(ctx context.Context, req *pb.Revenue) (*pb.Void, error) {
	k.Logger.Info("UpdateRevenue method is starting")

	err := k.Repo.UpdateRevenue(ctx, req.Id, req.Revenue)
	if err != nil {
		er := errors.Wrap(err, "failed to update revenue")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("UpdateRevenue has successfully finished")
	return &pb.Void{}, nil
}

func (k *KitchenService) GetName(ctx context.Context, req *pb.ID) (*pb.Name, error) {
	k.Logger.Info("GetName method is starting")

	name, err := k.Repo.GetName(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get name")
		k.Logger.Error(er.Error())
		return nil, er
	}

	k.Logger.Info("GetName has successfully finished")
	return &pb.Name{Name: name}, nil
}
