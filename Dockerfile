FROM golang:1.24 AS build

WORKDIR /src

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -o /geoip-auth \
    ./cmd/geoip-auth


FROM gcr.io/distroless/static-debian12

COPY --from=build /geoip-auth /geoip-auth

EXPOSE 8080

ENTRYPOINT ["/geoip-auth"]