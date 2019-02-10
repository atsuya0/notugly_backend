FROM golang:latest

RUN apt update \
    && apt install git

COPY ./go.mod ${root}/
COPY ./go.sum ${root}/

RUN go mod download
