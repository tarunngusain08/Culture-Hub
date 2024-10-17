FROM golang:1.22

WORKDIR /usr/src/app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go build -o /builds/main ./cmd/webapp/

EXPOSE 1323

ENTRYPOINT ["/builds/main"]
