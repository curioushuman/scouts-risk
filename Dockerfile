# syntax=docker/dockerfile:1

##
## Build the application from source
##

FROM golang:1.23 AS build-stage

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

RUN mkdir scouts-risk

COPY ./api/ ./scouts-risk/
COPY ./internal/ ./scouts-risk/
COPY ./web/ ./scouts-risk/

COPY ./cmd/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/main.go

##
## Deploy the application binary into a lean image
##

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./main"]