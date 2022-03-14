FROM golang:alpine AS build

RUN apk add --no-cache git openssh

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o bin/app main.go

FROM alpine:edge

RUN adduser -D -u 1000 junv

WORKDIR /home/junv

USER junv

COPY --from=build /app/bin/app .
COPY --from=build /app/test*.txt ./

CMD [ "/home/junv/app" ]
