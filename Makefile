ifeq ($(VERBOSE),1)
	V=''
else
	V=@
endif

include .env
.PHONY: docs

dev: mod swagger formatter vet test build

build: mod swagger
	@echo 
	@echo "*** building ***"
	@echo 
	$(V)go build

swagger:
	@echo 
	@echo "*** swagger definition ***"
	@echo 
	$(V)swag init -o docs
	
formatter:
	@echo 
	@echo "*** code formatting ***"
	@echo 
	$(V)go fmt ./...

vet:
	@echo 
	@echo  "*** code vetting ***"
	@echo 
	$(V)go vet ./...

test:
	@echo 
	@echo "*** unit testing ***"
	@echo 
	$(V)go test -v -race -buildvcs -coverprofile=coverage.out ./...
	$(V)go tool cover -func=coverage.out

mod:
	@echo 
	@echo "*** requirements ***"
	@echo 
	$(V)go mod tidy

dbu:
	$(V)migrate \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable \
		-source file://db/migrations \
		up

dbd:
	$(V)migrate \
		-database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)?sslmode=disable \
		-source file://db/migrations \
		down

infra-up:
	$(V)docker compose up

infra-zap:
	$(V)docker compose down