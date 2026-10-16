FROM golang:1.26.5-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/graphql ./graphql

FROM alpine:3.22
RUN apk --no-cache add ca-certificates
COPY --from=build /out/graphql /usr/local/bin/graphql
EXPOSE 8080
ENTRYPOINT ["graphql"]