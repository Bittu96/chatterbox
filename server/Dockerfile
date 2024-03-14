FROM golang:latest
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

ENV HTTP_PORT=8080
EXPOSE 8080

ENTRYPOINT ["/app/cb-staging"]