FROM golang:1.23.2-bullseye AS base

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backstagefood ./cmd/app/.

FROM alpine

WORKDIR /app

COPY --from=base /app/backstagefood .

CMD ["/app/backstagefood"]