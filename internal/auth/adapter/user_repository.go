package adapter

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"

	"zero/internal/auth/domain"
	"zero/internal/auth/domain/common"
)

type repoUser struct {
	ID        int       `db:"id"`
	UID       string    `db:"uid"`
	Email     string    `db:"email"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

const repoTableUser = "auth"

type repoColumnPatternUser struct {
	ID        string
	UID       string
	Email     string
	Name      string
	CreatedAt string
	UpdatedAt string
}

var repoColumnUser = repoColumnPatternUser{
	ID:        "id",
	UID:       "uid",
	Email:     "email",
	Name:      "name",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

func (c *repoColumnPatternUser) columns() string {
	return strings.Join([]string{
		c.ID,
		c.UID,
		c.Email,
		c.Name,
		c.CreatedAt,
		c.UpdatedAt,
	}, ", ")
}

func (r *PostgresRepository) CreateUser(ctx context.Context, param domain.User) (*domain.User, common.Error) {
	insert := map[string]interface{}{
		repoColumnUser.UID:   param.UID,
		repoColumnUser.Email: param.Email,
		repoColumnUser.Name:  param.Name,
	}
	// build SQL query
	query, args, err := r.pgsq.Insert(repoTableUser).
		SetMap(insert).
		Suffix(fmt.Sprintf("returning %s", repoColumnUser.columns())).
		ToSql()
	if err != nil {
		return nil, common.NewError(common.ErrorCodeInternalProcess, err)
	}

	// execute SQL query
	row := repoUser{}
	if err = r.db.GetContext(ctx, &row, query, args...); err != nil {
		return nil, common.NewError(common.ErrorCodeRemoteProcess, err)
	}

	// map the query result back to domain model
	user := domain.User(row)
	return &user, nil
}

func (r *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, common.Error) {
	query, args, err := r.pgsq.Select(repoColumnUser.columns()).
		From(repoTableUser).
		Where(sq.Eq{repoColumnUser.Email: email}).
		Limit(1).
		ToSql()
	if err != nil {
		return nil, common.NewError(common.ErrorCodeInternalProcess, err)
	}
	row := repoUser{}

	// get one row from result
	if err = r.db.GetContext(ctx, &row, query, args...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, common.NewError(common.ErrorCodeResourceNotFound, err, common.WithMsg("auth is not found"))
		}
		return nil, common.NewError(common.ErrorCodeRemoteProcess, err)
	}

	// map the query result back to domain model
	user := domain.User(row)
	return &user, nil
}

func (r *PostgresRepository) AuthenticateUser(_ context.Context, email string, password string) common.Error {
	// Authenticate the account

	return nil
}
