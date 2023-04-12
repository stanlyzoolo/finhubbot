# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

RUN apk add --no-cache build-base

WORKDIR /app

COPY . /app

RUN go env -w CGO_ENABLED='1'

RUN go mod download

RUN go build -o /finHubBot

CMD [ "/finHubBot" ]