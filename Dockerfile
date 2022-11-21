FROM golang:1.19-alpine AS build_base

#RUN apk add --no-cache git
RUN apk add  --no-cache build-base

# Set the Current Working Directory inside the container
WORKDIR /tmp/go-sample-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.* ./

RUN go mod tidy

COPY . .

# Unit tests
#RUN CGO_ENABLED=0 go test -v

# Build the Go app
RUN go build -race -o menu-sv

# Start fresh from a smaller image
FROM alpine:3.9

RUN apk add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /app

COPY --from=build_base /tmp/go-sample-app/menu-sv /app/go-app

# This container exposes port 8081 to the outside world
EXPOSE 50099

# Run the binary program produced by `go install`
CMD ["./go-app"]
