all: update docker 

update: 
	@go get -u ./...
	@go mod tidy

docker: 
	@docker build -t backstagefood:latest .
