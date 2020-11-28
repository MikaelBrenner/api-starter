FROM golang:1.15-buster

WORKDIR /server
COPY . .
EXPOSE 9797

CMD ["go", "run", "main.go"]

