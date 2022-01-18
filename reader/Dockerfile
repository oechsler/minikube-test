FROM golang:1.17-alpine as build

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY src .
RUN go build -o ./out/reader .

FROM alpine:latest

WORKDIR /opt

COPY --from=build /build/out/reader .

ENTRYPOINT [ "/opt/reader" ]