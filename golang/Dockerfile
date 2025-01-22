FROM golang:1.23 AS buildgo
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o ./build/tasks ./cmd/tasks

FROM alpine:3
WORKDIR /app
COPY --from=buildgo /app/build/tasks /app/tasks
ENTRYPOINT ["/app/tasks"]
