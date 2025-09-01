FROM golang:1.25rc2-alpine3.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app

RUN go install github.com/air-verse/air@latest

CMD ["air"]