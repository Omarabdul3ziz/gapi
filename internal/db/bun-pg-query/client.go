package bunpgquery

import (
	"context"
	"database/sql"

	"github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/db"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

var _ db.QueryClient = (*DBQueryClient)(nil)

type DBQueryClient struct {
	db *bun.DB
}

func NewDBQueryClient(dsn string) (*DBQueryClient, error) {
	pgdb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(pgdb, pgdialect.New())

	// debug
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return &DBQueryClient{
		db: db,
	}, nil
}

func (c *DBQueryClient) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

func (c *DBQueryClient) GetNodes(ctx context.Context) ([]db.Node, error) {
	var nodes []db.Node
	if err := c.db.NewSelect().
		TableExpr("node").
		ColumnExpr("node_id").
		Scan(ctx, &nodes); err != nil {
		return nil, err
	}

	return nodes, nil
}
