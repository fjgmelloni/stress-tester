FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go build -o stress-tester main.go

CMD ["./stress-tester"]