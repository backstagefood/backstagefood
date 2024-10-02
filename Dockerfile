FROM golang:1.23.2-bullseye AS base

WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go mod tidy \
&& go get -u \
&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backstagefood

FROM scratch

WORKDIR /app

COPY --from=base /app/backstagefood .

CMD ["/app/backstagefood"]