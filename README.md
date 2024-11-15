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

- lint
    `go fmt ./...`

- static / vulnerability check
    `go vet ./...`

- test
    `go test -v -race -buildvcs ./...`

- test coverage
	`go test -v -race -buildvcs -coverprofile=/tmp/coverage.out ./...`
	`go tool cover -html=/tmp/coverage.out`

- build
    `go build -o=/tmp/bin/${binary_name} ${main_package_path}`

- debug build

- logging

- env variables
    - IP
    - PORT