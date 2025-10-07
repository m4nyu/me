package main

import (
	"engineer/src/internal/components"
	"engineer/src/internal/handlers"
	"engineer/src/internal/middleware"
	"fmt"
	"log"
	"net/http"
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
	fileServer := http.FileServer(http.Dir("./static"))
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

	port := ":3000"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
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

