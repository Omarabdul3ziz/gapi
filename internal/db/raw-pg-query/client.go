package rawpgquery

import (
	"context"
	"database/sql"

	_ "embed"

	_ "github.com/lib/pq"

	"github.com/threefoldtech/tfgrid-sdk-go/gapi/internal/db"
)

var _ db.QueryClient = (*DBQueryClient)(nil)

type DBQueryClient struct {
	db *sql.DB
}

func NewDBQueryClient(dsn string) (*DBQueryClient, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return &DBQueryClient{
		db: db,
	}, nil
}

func (c *DBQueryClient) Ping(ctx context.Context) error {
	return c.db.PingContext(ctx)
}

//go:embed queries/get_nodes.sql
var getNodesQuery string

func (c *DBQueryClient) GetNodes(ctx context.Context) ([]db.Node, error) {
	rows, err := c.db.Query(getNodesQuery)
	if err != nil {
		return nil, err
	}

	var nodes []db.Node
	for rows.Next() {
		var node db.Node
		if err := rows.Scan(
			&node.NodeId,
		); err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}

func (c *DBQueryClient) GetFarms(ctx context.Context) []db.Farm {
	return nil
}

func (c *DBQueryClient) GetTwins(ctx context.Context) []db.Twin {
	return nil
}

func (c *DBQueryClient) GetContracts(ctx context.Context) []db.Contract {
	return nil
}
