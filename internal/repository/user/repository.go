package user

import (
	"context"
	"user-service/internal/model"
	"user-service/internal/repository"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	tableName         = `"user"`
	idColumn          = "id"
	firstNameColumn   = "first_name"
	lastNameColumn    = "last_name"
	passwordColumn    = "password"
	phoneNumberColumn = "phone_number"
	emailColumn       = "email"
	createdAtColumn   = "created_at"
	updatedAtColumn   = "updated_at"
)

var (
	psq = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, info *model.UserInfo) (int64, error) {

	builder := psq.Insert(tableName).
		Columns(firstNameColumn, lastNameColumn, passwordColumn, phoneNumberColumn, emailColumn).
		Values(info.FirstName, info.LastName, info.Password, info.PhoneNumber, info.Email).
		Suffix("RETURNING id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	var id int64
	err = r.db.QueryRow(ctx, query, args...).Scan(&id)
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

	var user model.User
	var info model.UserInfo
	err = r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &info.FirstName, &info.LastName, &info.Password, &info.PhoneNumber, &info.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}
	user.Info = &info
	return &user, nil
}
