FROM golang:1.19-alpine AS builder

RUN apk add --no-cache git make
RUN mkdir -p /app

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .


# Build the Go app
RUN make build

# Start fresh from a smaller image

FROM golang:1.19-alpine

COPY --from=builder /app/kaowao /bin/kaowao

WORKDIR /data

ENTRYPOINT ["kaowao"]
