FROM golang:latest

WORKDIR /app

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN ./migrate.sh
RUN go mod tidy && go build -o main .

EXPOSE 8000

CMD ["./main"]

