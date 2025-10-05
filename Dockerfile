# https://hub.docker.com/_/golang
FROM golang:1.25
WORKDIR /app

# Doing this before COPY . . is a cache optimization
COPY go.mod go.sum ./
RUN go mod download

# I think that this would be redundant since the fiber dependencies are listed in go.mod
#RUN go get github.com/gofiber/fiber/v2

COPY server.go .

RUN go build -o server 

