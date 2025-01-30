package user

import (
	"context"
	"user-service/internal/client/db"
	"user-service/internal/model"
	"user-service/internal/repository"

	sq "github.com/Masterminds/squirrel"
)

const (
	tableName         = `"user"`
	idColumn          = "id"
	firstNameColumn   = "first_name"
	lastNameColumn    = "last_name"
	passwordColumn    = "password"
	phoneNumberColumn = "phone_number"
	emailColumn       = "email"
	roleColumn        = "role"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

var (
	psq = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, info *model.UserInfo) (int64, error) {

	builder := psq.Insert(tableName).
		Columns(firstNameColumn, lastNameColumn, passwordColumn, phoneNumberColumn, emailColumn, roleColumn).
		Values(info.FirstName, info.LastName, info.Password, info.PhoneNumber, info.Email, info.Role).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRow: query,
	}
	var id int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := psq.Select(idColumn, firstNameColumn, lastNameColumn, passwordColumn, phoneNumberColumn, emailColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		Where(sq.Eq{idColumn: id}).
		Limit(1)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRow: query,
	}
	var user model.User
	err = r.db.DB().ScanOneContext(ctx, &user, q, args...)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
