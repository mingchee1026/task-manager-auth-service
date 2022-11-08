package postgres

import (
	"context"
	"fmt"
	"github.com/Levap123/task-manager-auth-service/proto"
	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}
func (ar *AuthRepo) Create(ctx context.Context, user *proto.User) (*proto.User, error) {
	tx, err := ar.db.Beginx()
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("INSER INTO %s(email, name, password) VALUES($1,$2,$3) RETURNING id", usersTable)
	if err := tx.GetContext(ctx, &user.Id, query, user.Email, user.Name, user.Password); err != nil {
		return nil, err
	}
	return user, nil
}
