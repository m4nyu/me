FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o server ./app/src/cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/server .
COPY --from=builder /build/app/src/static ./app/src/static

EXPOSE 3000

CMD ["./server"]
