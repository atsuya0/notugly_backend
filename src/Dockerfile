FROM golang:latest

RUN apt update \
    && apt install git

ARG ROOT
WORKDIR ${ROOT}

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download
