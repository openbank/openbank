# Builder
FROM golang:1.12.6-alpine3.9 as builder

RUN apk --no-cache add ca-certificates git

WORKDIR /src

# Copy go.mod and go.sum and run 'go mod download'.
# This way, docker will keep this costly download step cached as long as the
# go.mod and go.sum files don't change.
COPY go.* ./
RUN go mod download

COPY . .

ARG VERSION
RUN CGO_ENABLED=0 go build -ldflags="-w -s -X main.version=$VERSION" -o openbank

# Run the application
FROM alpine:3.9

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /src/openbank .
ADD env/sample.config env/config

CMD ["./openbank"]
