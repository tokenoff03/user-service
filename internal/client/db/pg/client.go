package pg

import (
	"context"
	"user-service/internal/client/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

type pgClient struct {
	masterDBC db.DB
}

func New(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &pgClient{
		masterDBC: &pg{dbc: dbc}, //Можно использовать конструктор в файле pg.go
	}, nil
}

func (c *pgClient) DB() db.DB {
	return c.masterDBC
}

func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
