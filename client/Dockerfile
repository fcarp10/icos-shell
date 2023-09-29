FROM golang:1.21 AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY pkg ./pkg
COPY cmd ./cmd
COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /shell-client

EXPOSE 8080
CMD ["/shell-client"]

