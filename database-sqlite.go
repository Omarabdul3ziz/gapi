package main

type SqliteDatabase struct{}

var _ Database = (*SqliteDatabase)(nil)

func NewSqliteDatabase() SqliteDatabase {
	return SqliteDatabase{}
}

func (db *SqliteDatabase) GetNodes() ([]Node, error) {
	return []Node{}, nil
}
