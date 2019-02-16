FROM golang:latest

RUN apt update \
    && apt install git

RUN go get golang.org/x/tools/cmd/goimports

ARG ROOT
WORKDIR ${ROOT}

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download
