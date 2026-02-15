FROM golang:1.25-alpine AS builder

WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o termfolio .

FROM alpine:3.21

WORKDIR /app
COPY --from=builder /build/termfolio .
RUN mkdir -p /app/data

EXPOSE 2222

ENTRYPOINT ["./termfolio", "--serve", "--port", "2222", "--key", "/app/data/hostkey"]
