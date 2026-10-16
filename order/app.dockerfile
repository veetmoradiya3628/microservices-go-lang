FROM golang:1.26.5-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/order ./order/cmd/order

FROM alpine:3.22
RUN apk --no-cache add ca-certificates
COPY --from=build /out/order /usr/local/bin/order
EXPOSE 8080
ENTRYPOINT ["order"]