FROM golang:1.22.7-bullseye AS builder
WORKDIR /app
COPY .. .
RUN go mod download
RUN go build -o /app/sso /app/cmd/sso/main.go
RUN go build -o /app/migrator /app/cmd/migrator/main.go

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/sso /app/sso
COPY --from=builder /app/migrator /app/migrator
COPY ./migrations ./migrations
ENV COMMAND=app

ENTRYPOINT ["/bin/sh", "-c", "if [ \"$COMMAND\" = 'migrator' ]; then /app/migrator; else /app/sso; fi"]
