package repository

import (
	"context"
	"database/sql"
	"errors"

	"go-chi-template/internal/app/user"
	"go-chi-template/internal/pkg/db"

	"github.com/google/uuid"
)

type UserRepository struct {
	db db.SQL
}

var _ user.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db db.SQL) *UserRepository {
	return &UserRepository{db: db}
}

const queryUpsertUser = `
	INSERT INTO users (id, full_name)
	VALUES ($1, $2)
	ON CONFLICT(id) DO UPDATE
	SET full_name = :full_name
	RETURNING id, full_name
`

func (u UserRepository) UpsertUser(ctx context.Context, userParam user.User) (res user.User, err error) {
	rows, err := u.db.QueryContext(ctx, queryUpsertUser, userParam.ID, userParam.FullName)
	if err != nil {
		return res, err
	}
	defer func() { _ = rows.Close() }()

	if rows.Next() {
		err = rows.Scan(&res.ID, &res.FullName)
		if err != nil {
			return res, err
		}
	}

	return res, nil
}

const queryGetUser = `
	SELECT id, full_name
	FROM users
	WHERE id = $1
`

func (u UserRepository) GetUser(ctx context.Context, userID uuid.UUID) (*user.User, error) {
	rows, err := u.db.QueryContext(ctx, queryGetUser, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	if rows.Next() {
		var res user.User
		err = rows.Scan(&res.ID, &res.FullName)
		if err != nil {
			return nil, err
		}
		return &res, nil
	}

	return nil, nil
}
