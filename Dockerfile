FROM golang:alpine

RUN apk add --no-cache git openssh

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod download

CMD ["go", "run", "main.go" ]
