package main

import (
	"engineer/src/app/internal/components"
	"engineer/src/app/internal/handlers"
	"engineer/src/app/internal/middleware"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Middleware
	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CookieMiddleware)

	// Static files
	fileServer := http.FileServer(http.Dir("./src/app/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	// Language routes - explicit routes for each language
	for _, lang := range []string{"EN", "DE", "FR", "ES", "IT", "PT", "NL", "RU", "JA", "KO", "ZH", "AR"} {
		langCode := lang // capture for closure
		r.Get("/"+strings.ToLower(langCode), func(w http.ResponseWriter, r *http.Request) {
			handleHomeWithLang(w, r, langCode)
		})
	}

	// Root redirect to default language
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		lang := middleware.GetLang(r)
		http.Redirect(w, r, "/"+strings.ToLower(lang), http.StatusTemporaryRedirect)
	})

	// API routes for setting preferences
	r.Post("/api/theme", handlers.SetTheme)
	r.Post("/api/lang", handlers.SetLanguage)

	// Live reload WebSocket endpoint (development only)
	r.Get("/livereload", middleware.LiveReloadHandler)

	port := getPort()
	fmt.Printf("Server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

// getPort returns an available port, starting from 3000
func getPort() string {
	// Check if PORT env var is set (production)
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}

	// For local dev, try ports starting from 3000
	startPort := 3000
	for port := startPort; port < startPort+10; port++ {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			listener.Close()
			return addr
		}
	}

	// Fallback to 3000 if all ports are taken
	return ":3000"
}

// handleHomeWithLang handles homepage with a specific language code
func handleHomeWithLang(w http.ResponseWriter, r *http.Request, langCode string) {
	theme := middleware.GetTheme(r)

	page := components.BaseLayout("4nuel", theme,
		components.HomePage(components.HomePageProps{
			Theme:       theme,
			CurrentLang: langCode,
		}),
	)

	w.Header().Set("Content-Type", "text/html")
	_ = page.Render(w)
}

