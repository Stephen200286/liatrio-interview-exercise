# https://hub.docker.com/_/golang
FROM golang:1.25
WORKDIR /app

COPY go.mod go.sum server.go ./
RUN go mod download

#COPY server.go .
EXPOSE 80

RUN go build -o server 

CMD ["./server"]
