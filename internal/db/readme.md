# Database

what ever happens here, the external api shouldn't know about, orm could be changed or even using a raw sql. but all should implement a client that implement the interface in db.go and expose the types in models.go

## Decisions

- raw sql is the best in performance
- bun don't hide sql, and help with scanning to types

## Notes

- what if you gonna need models? it will differ between the orms
- add limits/order as a middleware method
- how to handle filter
