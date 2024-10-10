all: update docker 

update: 
	@go get -u ./...
	@go mod tidy

build: 
	@docker build -t backstagefood:latest .

up:
	@docker compose up --build