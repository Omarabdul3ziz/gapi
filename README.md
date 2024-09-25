# G-API

TF Grid-API provides all the data clients need from chain and grid.

## Structure

Database:

- QueryClient: read-only-access to the db
- MutationClient: write-access to the db
- ChainDb: a dump of chain events
- GridDb: normalized/de-normalized tables

Services:

- Processor: process data from chain-db > grid-db (mutation)
- Indexer: process data from node > grid-db (mutation)

API:

- APIClient: interface invokes the (db-query-client)
- Rest: implement a rest api server
- GraphQl: implement a graphql api server
- RPC: implement an RPC api server

## Mindset

- modular: its is a building blocks, each component can be replaced with another doing the same function.
- fast: performance is number one concern
- extendable: very easy to go through and add extra features.

## External tools and why

- Gorm: best orm out there
- sqlc: generate from sql
- gorilla/schema: parse params
- gorilla/mux: serve rest api

## Should consider

- proper schema which have foreign keys
- postgres views the raw data
- a well defined documentation detailed on each topic
- think about db cluster with cockroachdb
- add indexes
- consider caching the most frequent queries, maybe with `no-cache` query param
- abstract the layers. dbclient(query/mutate) shouldn't know what orm/raw sql it uses.
- api versioning
- start rest/graphql/processor/indexer/griddb/chaindb from separate binaries

## Decisions

- consider having two databases, one that dumps chain event and the other are processed from it into denormalized tables.
- the grid db should have two accesses one that can mutate, only works internally for processor/indexer and the other is read only can be exposed to test a running instance for example. and used on the server part.

## General Notes

`internal/db` have a various types of clients. a client can be identified with `<orm>-<dialect>-<operation>` each of these clients should implement the interface in `internal/db/db.go`. now we have

- `raw-pg-query` client: which is a postgres query client uses raw sql statements.
- `gorm-pg-query` client: uses gorm
- `bun-pg-query` client: uses bun as an orm

## Phases

- the Rest/Graphql api
- the database query client
- the schema building
- the trigger mutations
- the processor mutations
- the indexer
