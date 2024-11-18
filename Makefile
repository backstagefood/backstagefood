VERSION := $(shell cat VERSION)

all: update up

run: swagger exec

exec:
	@go run cmd/app/main.go

update: 
	@go get -u ./...
	@go mod tidy

docker-build:
	@docker build -t backstagefood:$(VERSION) .

setup-env:
	@cp .env.example .env

up:
	@if [ ! -f .env ]; then $(MAKE) setup-env; fi
	@docker compose up --build

swagger:
# install -> go install github.com/swaggo/swag/cmd/swag@latest
# install -> go get github.com/swaggo/echo-swagger@latest
	swag init -g cmd/app/main.go