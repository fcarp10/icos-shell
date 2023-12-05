FROM golang:1.21 AS build

LABEL maintainer="Marc Michalke"

WORKDIR /app

COPY go.* ./
COPY *.go ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /registration

EXPOSE 8080
CMD ["/registration"]

