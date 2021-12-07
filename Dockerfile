FROM golang:1.17.4 AS build-env

RUN mkdir /users-service
WORKDIR /users-service
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /users-server ./cmd/userssrv/main.go

FROM alpine:latest

EXPOSE 8080

WORKDIR /
COPY --from=build-env /users-server /
RUN chmod 755 /users-server

CMD ["/users-server"]