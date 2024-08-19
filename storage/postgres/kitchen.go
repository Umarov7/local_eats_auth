package postgres

import (
	pb "auth-service/genproto/kitchen"
	"context"
	"database/sql"
	"log"
	"strconv"

	"github.com/pkg/errors"
)

type KitchenRepo struct {
	DB *sql.DB
}

func NewKitchenRepo(db *sql.DB) *KitchenRepo {
	return &KitchenRepo{DB: db}
}

func (k *KitchenRepo) Create(ctx context.Context, data *pb.CreateRequest) (*pb.CreateResponse, error) {
	query := `
	insert into
		kitchens (owner_id, name, description, cuisine_type, address, phone_number)
	values
		($1, $2, $3, $4, $5, $6)
	returning
		id, owner_id, name, description, cuisine_type, address, phone_number, rating, created_at
	`

	row := k.DB.QueryRowContext(ctx, query, data.OwnerId, data.Name, data.Description,
		data.CuisineType, data.Address, data.PhoneNumber)

	var kn pb.CreateResponse

	err := row.Scan(&kn.Id, &kn.OwnerId, &kn.Name, &kn.Description, &kn.CuisineType,
		&kn.Address, &kn.PhoneNumber, &kn.Rating, &kn.CreatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "insertion failure")
	}

	return &kn, nil
}

func (k *KitchenRepo) Read(ctx context.Context, id *pb.ID) (*pb.Info, error) {
	query := `
	select
		owner_id, name, description, cuisine_type, address,
		phone_number, rating, total_orders, created_at, updated_at 
	from
		kitchens
	where
		deleted_at is null and id = $1
	`

	kn := pb.Info{Id: id.Id}

	err := k.DB.QueryRowContext(ctx, query, kn.Id).Scan(&kn.OwnerId, &kn.Name, &kn.Description,
		&kn.CuisineType, &kn.Address, &kn.PhoneNumber, &kn.Rating, &kn.TotalOrders,
		&kn.CreatedAt, &kn.UpdatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "reading failure")
	}

	return &kn, nil
}

func (k *KitchenRepo) Update(ctx context.Context, data *pb.NewData) (*pb.UpdatedData, error) {
	tr, err := k.DB.Begin()
	if err != nil {
		return nil, errors.Wrap(err, "transaction failure")
	}

	defer func() {
		if err != nil {
			if rbErr := tr.Rollback(); rbErr != nil {
				log.Println("error rolling back transaction:", rbErr)
			}
		} else {
			if cErr := tr.Commit(); cErr != nil {
				log.Println("error committing transaction:", cErr)
			}
		}
	}()

	query := `
	update
		kitchens
	set
		name = $1, description = $2, phone_number = $3, updated_at = NOW()
	where
		deleted_at is null and id = $4
	`

	res, err := tr.ExecContext(ctx, query, data.Name, data.Description, data.PhoneNumber, data.Id)
	if err != nil {
		return nil, errors.Wrap(err, "update failure")
	}

	rowsNum, err := res.RowsAffected()
	if err != nil {
		return nil, errors.Wrap(err, "rows affected failure")
	}
	if rowsNum < 1 {
		return nil, errors.Wrap(err, "no rows affected")
	}

	query = `
	select
		owner_id, name, description, cuisine_type,
		address, phone_number, rating, updated_at
	from
		kitchens
	where
		deleted_at is null and id = $1
	`

	upData := pb.UpdatedData{Id: data.Id}

	err = tr.QueryRowContext(ctx, query, upData.Id).Scan(&upData.OwnerId, &upData.Name,
		&upData.Description, &upData.CuisineType, &upData.Address,
		&upData.PhoneNumber, &upData.Rating, &upData.UpdatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "reading failure")
	}

	return &upData, nil
}

func (k *KitchenRepo) Delete(ctx context.Context, id *pb.ID) error {
	query := `
	update
		kitchens
	set
		deleted_at = NOW()
	where
		deleted_at is null and id = $1
	`

	res, err := k.DB.ExecContext(ctx, query, id.Id)
	if err != nil {
		return errors.Wrap(err, "deletion failure")
	}

	rowsNum, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "rows affected failure")
	}
	if rowsNum < 1 {
		return errors.Wrap(err, "no rows affected")
	}

	return nil
}

func (k *KitchenRepo) Fetch(ctx context.Context, pag *pb.Pagination) ([]*pb.KitchenDetails, int, error) {
	query := `
	select
		id, name, cuisine_type, rating, total_orders
	from
		kitchens
	where
		deleted_at is null
	limit $1
	offset $2
	`

	rows, err := k.DB.QueryContext(ctx, query, pag.Limit, pag.Offset)
	if err != nil {
		return nil, -1, errors.Wrap(err, "retrieval failure")
	}
	defer rows.Close()

	var kitchens []*pb.KitchenDetails
	for rows.Next() {
		var kn pb.KitchenDetails

		err := rows.Scan(&kn.Id, &kn.Name, &kn.CuisineType, &kn.Rating, &kn.TotalOrders)
		if err != nil {
			return nil, -1, errors.Wrap(err, "reading failure")
		}

		kitchens = append(kitchens, &kn)
	}

	rowsNum, err := k.CountRows(ctx)
	if err != nil {
		return nil, -1, err
	}

	return kitchens, rowsNum, nil
}

func (k *KitchenRepo) Search(ctx context.Context, det *pb.SearchDetails) ([]*pb.KitchenDetails, int, error) {
	query := `
	select
		id, name, cuisine_type, rating, total_orders
	from
		kitchens
	where
		deleted_at is null
	`

	var params []interface{}

	if det.Query != "" {
		query += " and (name ILIKE $1 or description ILIKE $1)"
		params = append(params, "%"+det.Query+"%")
	}
	if det.CuisineType != "" {
		query += " and cuisine_type = $" + strconv.Itoa(len(params)+1)
		params = append(params, det.CuisineType)
	}
	if det.Rating > 0 {
		query += " and rating >= $" + strconv.Itoa(len(params)+1)
		params = append(params, det.Rating)
	}
	if det.Pagination.Limit > 0 {
		query += " limit $" + strconv.Itoa(len(params)+1)
		params = append(params, det.Pagination.Limit)
	}
	if det.Pagination.Offset > 0 {
		query += " offset $" + strconv.Itoa(len(params)+1)
		params = append(params, det.Pagination.Offset)
	}

	rows, err := k.DB.QueryContext(ctx, query, params...)
	if err != nil {
		return nil, -1, errors.Wrap(err, "retrieval failure")
	}
	defer rows.Close()

	var kitchens []*pb.KitchenDetails
	for rows.Next() {
		var kn pb.KitchenDetails

		err := rows.Scan(&kn.Id, &kn.Name, &kn.CuisineType, &kn.Rating, &kn.TotalOrders)
		if err != nil {
			return nil, -1, errors.Wrap(err, "reading failure")
		}

		kitchens = append(kitchens, &kn)
	}

	rowsNum, err := k.CountRows(ctx)
	if err != nil {
		return nil, -1, err
	}

	return kitchens, rowsNum, nil
}

func (k *KitchenRepo) CountRows(ctx context.Context) (int, error) {
	var rowsNum int
	query := "select count(1) from kitchens where deleted_at is null"

	err := k.DB.QueryRowContext(ctx, query).Scan(&rowsNum)
	if err != nil {
		return -1, errors.Wrap(err, "rows counting failure")
	}

	return rowsNum, nil
}

func (k *KitchenRepo) ValidateKitchen(ctx context.Context, id string) (bool, error) {
	query := `
    select EXISTS (
		select 1
		from kitchens
		where deleted_at is null and id = $1
	)`

	var status bool
	err := k.DB.QueryRowContext(ctx, query, id).Scan(&status)
	if err != nil {
		return false, errors.Wrap(err, "validation failure")
	}

	if !status {
		return false, errors.New("kitchen not found")
	}

	return status, nil
}

func (k *KitchenRepo) IncrementTotalOrders(ctx context.Context, id string) error {
	query := `
	update
		kitchens
	set
		total_orders = total_orders + 1
	where
		deleted_at is null and id = $1
	`

	_, err := k.DB.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "increment failure")
	}

	return nil
}

func (k *KitchenRepo) UpdateRating(ctx context.Context, id string, rating float32) error {
	query := `
	update
		kitchens
	set
		rating = (rating + $1) / 2
	where
		deleted_at is null and id = $2
	`

	_, err := k.DB.ExecContext(ctx, query, rating, id)
	if err != nil {
		return errors.Wrap(err, "update failure")
	}

	return nil
}

func (k *KitchenRepo) UpdateRevenue(ctx context.Context, id string, revenue float32) error {
	query := `
	update
		kitchens
	set
		total_revenue = total_revenue + $1
	where
		deleted_at is null and id = $2
	`

	_, err := k.DB.ExecContext(ctx, query, revenue, id)
	if err != nil {
		return errors.Wrap(err, "revenue update failure")
	}

	return nil
}

func (k *KitchenRepo) GetName(ctx context.Context, id string) (string, error) {
	query := `
	select
		name
	from
		kitchens
	where deleted_at is null and id = $1
	`

	var name string
	err := k.DB.QueryRowContext(ctx, query, id).Scan(&name)
	if err != nil {
		return "", errors.Wrap(err, "name retrieval failure")
	}

	return name, nil
}
