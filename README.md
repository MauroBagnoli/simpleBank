# Project initialization

- To fix errors after query creation

update go:
go get -u ./...

install pgtype
go get github.com/jackc/pgx/v5/pgtype
go clean -modcache
