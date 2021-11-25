FROM golang:1.14.9-alpine
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build
