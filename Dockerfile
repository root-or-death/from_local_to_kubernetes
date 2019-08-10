FROM golang:1.11

COPY . .

EXPOSE 80

CMD go run main.go
