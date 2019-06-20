# Builder

FROM golang:1.12.6-alpine3.9 as builder

WORKDIR /go/src/github.com/openbank/openbank

COPY . .

ARG VERSION
RUN CGO_ENABLE=0 go build -ldflags="-w -s -X main.version=$VERSION" -o openbank

# Run the application
FROM alpine:3.9

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/
COPY --from=builder /go/src/github.com/openbank/openbank/openbank .
ADD env/sample.config env/config

CMD ["./openbank"]
