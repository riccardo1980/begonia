all: begonia
.PHONY: docs

begonia: docs mod
	go build

docs:
	# go get -u github.com/swaggo/gin-swagger
	# go get -u github.com/swaggo/files
	swag init -g main.go -o docs

mod:
	go mod tidy