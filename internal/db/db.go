package db

import "context"

type QueryClient interface {
	Ping(context.Context) error
	GetNodes(context.Context) ([]Node, error)
	// GetFarms(context.Context) []Farm
	// GetTwins(context.Context) []Twin
	// GetContracts(context.Context) []Contract
}

type MutationClient interface{}
