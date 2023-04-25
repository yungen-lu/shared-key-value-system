# modules caching
FROM golang:1.20-bullseye as modules

COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

# builder
FROM golang:1.20-bullseye as build

COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o /go/bin/shared-key-value-list-system ./app

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/shared-key-value-list-system /
# ENTRYPOINT [ "/app-chat" ]
CMD ["/shared-key-value-list-system"]
