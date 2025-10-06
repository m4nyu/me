# Personal Website - Go Implementation

A pure Go implementation using gomponents for server-side rendering with no JavaScript (except Tailwind CDN).

## Prerequisites

- Go 1.24+
- Air (for hot reloading): `go install github.com/air-verse/air@v1.52.3`

## Development

Start the development server with hot reloading (watches for changes in `.go`, `.css`, `.html` files):

```bash
make
```

or

```bash
make dev
```

The server will start on http://localhost:3000 and automatically reload when you make changes.

## Other Commands

```bash
make build    # Build the application
make run      # Build and run without hot reloading
make clean    # Clean build artifacts
make help     # Show all available commands
```

## Project Structure

```
.
├── src/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go          # Entry point
│   └── internal/
│       ├── components/          # Gomponents (HTML components)
│       │   ├── accordion.go
│       │   ├── homepage.go
│       │   ├── lang.go
│       │   ├── layout.go
│       │   └── logo.go
│       ├── handlers/            # HTTP handlers
│       │   └── cookies.go
│       └── middleware/          # HTTP middleware
│           └── middleware.go
├── static/
│   ├── css/
│   │   └── globals.css      # Custom CSS (Tailwind via CDN)
│   ├── dark.svg
│   └── light.svg
├── .air.toml                # Hot reload configuration
└── Makefile                 # Build commands
```

## Features

- **Pure CSS interactivity** (no JavaScript)
  - Accordion using `<details>/<summary>`
  - Modal using `:target` pseudo-class
  - Carousel using radio buttons and `:checked`
- **Server-side rendering** with gomponents
- **Cookie-based state** for theme and language preferences
- **Hot reloading** for development
- **Tailwind CSS** via CDN
