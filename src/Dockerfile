FROM golang:latest AS develop
RUN go get golang.org/x/tools/cmd/goimports
ARG ROOT
WORKDIR ${ROOT}
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

FROM golang:latest AS build
COPY ./ /app
RUN go build -o app /app
CMD /app/app

FROM golang:alpine AS run
COPY --from=build /app/app /usr/local/bin/app
CMD /usr/local/bin/app
