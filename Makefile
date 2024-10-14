all: update up

update: 
	@go get -u ./...
	@go mod tidy

build: 
	@docker build -t backstagefood:latest .

up:
	@docker compose up --build

swagger:
# install -> go install github.com/swaggo/swag/cmd/swag@latest
# install -> go get github.com/swaggo/echo-swagger@latest
	swag init -g cmd/app/main.go