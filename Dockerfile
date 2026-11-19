# Static Go binary on distroless: non-root, no shell, no package manager.
FROM golang:1.27.4-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /server ./src/app/cmd/server

FROM gcr.io/distroless/static-debian13:nonroot
WORKDIR /app
COPY --from=build /server ./server
COPY --from=build /src/src/app/static ./src/app/static
ENV PORT=3000
EXPOSE 3000
ENTRYPOINT ["/app/server"]
