FROM golang:1.20-alpine AS build
RUN apk --no-cache add gcc g++ ca-certificates
WORKDIR /go/src/github.com/veetmoradiya3628/microservices-go-lang
COPY go.mod go.sum ./
COPY vendor vendor
COPY catalog catalog
RUN GO111MODULE=on go build -mod vendor -o /go/bin/app ./catalog/cmd/catalog

FROM alpine:3.18
WORKDIR /usr/bin
COPY --from=build /go/bin .
EXPOSE 8080
CMD ["./app"]