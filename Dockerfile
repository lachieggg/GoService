FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY src/ .

RUN go build -o goservice .

CMD ["./goservice"]