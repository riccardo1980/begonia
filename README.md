- B/E framework: gin
- ORM: gorm
- migration tool: goose
- swagger generation: go-swagger (gin?)


To deepen:
`go mod tidy -diff`
`go mod verify`
`go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...`
    `go run golang.org/x/vuln/cmd/govulncheck@latest ./...`

- module prerequisites update


- logging

- env variables
    - IP
    - PORT


# golang-migrate
```BASH
 go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
- HowTo: https://blog.didiktrisusanto.dev/how-to-use-golang-migrate-for-database-migration

`

migrate create -ext sql -dir db/migrations -seq create_users_table
# fill up and down files

# procede with migrations
export POSTGRESQL_URL=postgres://postgres:mysecretpassword@localhost:5432/example?sslmode=disable
migrate -database ${POSTGRESQL_URL} -source file://db/migrations up
`