package main

type Database interface {
	GetNodes() ([]Node, error)
	// GetNode(uint32) (Node, error)
	// AddNode(Node) error
	// DeleteNode(uint32) error
	// UpdateNode(Node) (Node, error)
}

type Node struct {
	NodeId uint32
}
