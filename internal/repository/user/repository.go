package user

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/kom1ssar/grpc-auth/internal/client/db"
	"github.com/kom1ssar/grpc-auth/internal/model"
	desc "github.com/kom1ssar/grpc-auth/internal/repository"
	"github.com/kom1ssar/grpc-auth/internal/repository/user/converter"
	modelRepo "github.com/kom1ssar/grpc-auth/internal/repository/user/model"
)

var _ desc.UserRepository = (*repository)(nil)

const (
	tableName       = "users"
	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	roleColumn      = "role"
	passwordColumn  = "password"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repository struct {
	db db.Client
}

func NewUserRepository(db db.Client) desc.UserRepository {
	return &repository{db: db}
}

func (r *repository) GetById(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(
		idColumn, nameColumn, emailColumn, roleColumn,
		passwordColumn, createdAtColumn, updatedAtColumn,
	).
		PlaceholderFormat(sq.Dollar).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.GetById",
		QueryRaw: query,
	}

	var user modelRepo.User

	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repository) Create(ctx context.Context, user *model.User) (int64, error) {
	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn).
		Values(user.Name, user.Email, user.Password).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var id int64

	err = r.db.DB().ScanOneContext(ctx, &id, q, args...)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *repository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn,
		passwordColumn, createdAtColumn, updatedAtColumn,
	).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{emailColumn: email}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var user modelRepo.User

	q := db.Query{
		Name:     "user_repository",
		QueryRaw: query,
	}

	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repository) GetByName(ctx context.Context, name string) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn,
		passwordColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{nameColumn: name}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository",
		QueryRaw: query,
	}

	var user modelRepo.User

	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)
	if err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}
