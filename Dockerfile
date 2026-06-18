FROM golang:1.25.8-alpine

WORKDIR /go/src/terraform-provider-haproxy
COPY . .

RUN go mod download
RUN go install -v ./...