FROM golang:1.16.3-alpine3.13

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
