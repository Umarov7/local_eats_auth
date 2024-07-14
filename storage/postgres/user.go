package postgres

import (
	pba "auth-service/genproto/auth"
	pb "auth-service/genproto/user"
	"context"
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) Create(ctx context.Context, info *pba.RegisterRequest) (*pba.RegisterResponse, error) {
	query := `
	insert into
		users (username, email, password_hash, full_name, user_type)
	values
		($1, $2, $3, $4, $5)
	returning
		id, username, email, full_name, user_type, created_at
	`

	var user pba.RegisterResponse
	row := u.DB.QueryRowContext(ctx, query,
		info.Username, info.Email, info.Password, info.FullName, info.UserType,
	)
	err := row.Scan(&user.Id, &user.Username, &user.Email,
		&user.FullName, &user.UserType, &user.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "insertion/scanning failure")
	}

	return &user, nil
}

func (u *UserRepo) Read(ctx context.Context, id *pb.ID) (*pb.Profile, error) {
	query := `
	select
		username, email, full_name, user_type, address, phone_number, created_at, updated_at
	from
		users
	where
		deleted_at is null and id = $1
	`

	pr := pb.Profile{Id: id.Id}

	err := u.DB.QueryRowContext(ctx, query, pr.Id).Scan(&pr.Username, &pr.Email, &pr.FullName,
		&pr.UserType, &pr.Address, &pr.PhoneNumber, &pr.CreatedAt, &pr.UpdatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "reading failure")
	}

	return &pr, nil
}

func (u *UserRepo) Update(ctx context.Context, data *pb.NewInfo) (*pb.Details, error) {
	tr, err := u.DB.Begin()
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
		users
	set
		full_name = $1, address = $2, phone_number = $3, updated_at = NOW()
	where
		deleted_at is null and id = $4
	`

	res, err := tr.ExecContext(ctx, query, data.FullName, data.Address, data.PhoneNumber, data.Id)
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
		username, email, full_name, user_type, address, phone_number, updated_at
	from
		users
	where
		deleted_at is null and id = $1
	`

	d := pb.Details{Id: data.Id}

	err = tr.QueryRowContext(ctx, query, d.Id).Scan(&d.Username, &d.Email, &d.FullName,
		&d.UserType, &d.Address, &d.PhoneNumber, &d.UpdatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "reading failure")
	}

	return &d, nil
}

func (u *UserRepo) Delete(ctx context.Context, id *pb.ID) error {
	tr, err := u.DB.Begin()
	if err != nil {
		return errors.Wrap(err, "transaction failure")
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
		users
	set
		deleted_at = NOW()
	where
		deleted_at is null and id = $1
	`

	res, err := tr.ExecContext(ctx, query, id.Id)
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

	query = `
	update
		kitchens
	set
		deleted_at = NOW()
	where
		deleted_at is null and owner_id = $1
	`

	_, err = tr.ExecContext(ctx, query, id.Id)
	if err != nil {
		return errors.Wrap(err, "owned kitchens' deletion failure")
	}

	return nil
}
