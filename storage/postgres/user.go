package postgres

import (
	"context"
	"database/sql"

	"github.com/Sharofiddin701/medium-clone-simple/storage/repo"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) repo.UserStorageI {

	return &UserRepo{
		db: db,
	}

}
func (u *UserRepo) Create(ctx context.Context, req *repo.User) (*repo.User, error) {
	query := ` 
    INSERT INTO users (
        id,
        first_name,
        last_name,
        email,
        password
    ) VALUES($1,$2,$3,$4,$5) RETURNING created_at
    `
	err := u.db.QueryRow(query, req.ID, req.FirstName, req.LastName, req.Email, req.Password).Scan(&req.CreatedAt)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (u *UserRepo) Update(ctx context.Context, req *repo.UpdateUser) error {
	tsx, err := u.db.Begin()

	if err != nil {
		return err
	}

	query := `
    UPDATE users SET
        first_name=$1,
        last_name=$2
    WHERE id=$3
    `
	res, err := tsx.Exec(query, req.FirstName, req.LastName, req.ID)

	if err != nil {
		errRoll := tsx.Rollback()
		if errRoll != nil {
			err = errRoll
		}
		return err
	}

	data, err := res.RowsAffected()
	if err != nil {
		errRoll := tsx.Rollback()
		if errRoll != nil {
			err = errRoll
		}
		return err
	}

	if data == 0 {
		tsx.Commit()
		return sql.ErrNoRows
	}
	return tsx.Commit()
}

func (u *UserRepo) Get(ctx context.Context, id string) (*repo.User, error) {
	query := `
	SELECT 
		id,
		first_name,
		last_name,
		email,
		password,
		created_at
		FROM users  WHERE id=$1
	`
	var user repo.User
	err := u.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return &user, err
	}
	return nil, nil
}

func (u *UserRepo) Delete(ctx context.Context, id string) error {
	tsx, err := u.db.Begin()

	if err != nil {
		return err
	}

	res, err := tsx.Exec("DELETE FROM users WHERE id =$1", id)
	if err != nil {
		errRoll := tsx.Rollback()
		if errRoll != nil {
			err = errRoll
		}
		return err
	}
	data, err := res.RowsAffected()
	if err != nil {
		errRoll := tsx.Rollback()
		if errRoll != nil {
			err = errRoll
		}
		return err
	}

	if data == 0 {
		tsx.Commit()
		return sql.ErrNoRows
	}

	return tsx.Commit()
}
