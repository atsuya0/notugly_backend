FROM golang:latest

RUN go get golang.org/x/tools/cmd/goimports

ARG ROOT
WORKDIR ${ROOT}

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download
