# Static Go binary on distroless: non-root, no shell, no package manager.
# CSS is compiled at build time; nothing compiles styles in the browser.
FROM node:26-slim AS css
WORKDIR /src
# Versions come from package-lock.json, which Dependabot keeps current.
COPY package.json package-lock.json ./
RUN npm ci --no-audit --no-fund
COPY src/app ./src/app
RUN npx @tailwindcss/cli -i src/app/static/css/input.css -o /app.css --minify

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
COPY --from=css /app.css ./src/app/static/css/app.css
ENV PORT=3000
EXPOSE 3000
ENTRYPOINT ["/app/server"]
