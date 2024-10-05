##
## Build
##

FROM golang:1.23 AS build
# FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

# RUN go build -o /main ./main.go
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /main ./main.go

##
## Deploy
##

FROM gcr.io/distroless/base-debian10
# FROM alpine:latest

WORKDIR /

COPY --from=build /main /main
# RUN chmod a+x /main

EXPOSE 8080

# this is debian only
USER nonroot:nonroot

ENTRYPOINT ["/main"]
