package gormpgquery

import (
	"context"

	"github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ db.QueryClient = (*DBQueryClient)(nil)

type DBQueryClient struct {
	db *gorm.DB
}

func NewDBQueryClient(dsn string) (*DBQueryClient, error) {
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	return &DBQueryClient{
		db: db,
	}, nil
}

func (c *DBQueryClient) Ping(ctx context.Context) error {
	db, err := c.db.DB()
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}

func (c *DBQueryClient) GetNodes(ctx context.Context) ([]db.Node, error) {
	var nodes []db.Node
	if err := c.db.Table("node").
		Select("node_id").
		Scan(&nodes); err.Error != nil {
		return nil, err.Error
	}

	return nodes, nil
}
