# Fetch
FROM golang:1.23-alpine AS fetch-stage
COPY go.mod go.sum /app
WORKDIR /app
RUN go mod download

# Generate
FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app
RUN ["templ", "generate"]

# Build
FROM golang:1.23-alpine AS build-stage
COPY --from=generate-stage /app /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app ./cmd/main.go

# Test
FROM build-stage AS test-stage
RUN go test -v ./...

# Deploy
FROM alpine:latest
WORKDIR /
COPY --from=build-stage /app/app /app
RUN chmod a+x /app

# Expose
EXPOSE 8080
ENTRYPOINT ["/app"]