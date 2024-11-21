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