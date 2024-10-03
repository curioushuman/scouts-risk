# syntax=docker/dockerfile:1

##
## Build the application from source
##

# FROM golang:1.23
FROM golang:1.23-alpine AS build-stage

# RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./

# RUN go mod download && go mod verify
RUN go mod download

COPY . .

# RUN CGO_ENABLED=0 GOOS=linux go build -o /main
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o /main

# RUN go build -v -o ./main.go ./app
# RUN chmod +x ./main

##
## Deploy the application binary into a lean image
##

# FROM gcr.io/distroless/base-debian11 AS build-release-stage
FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /main /main
RUN chmod a+x /main

EXPOSE 8080

CMD ["/main"]
