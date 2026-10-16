FROM golang:1.26.5-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/account ./account/cmd/account

FROM alpine:3.22
RUN apk --no-cache add ca-certificates
COPY --from=build /out/account /usr/local/bin/account
EXPOSE 8080
ENTRYPOINT ["account"]