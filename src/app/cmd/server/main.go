package main

import (
	"engineer/src/app/internal/components"
	"engineer/src/app/internal/handlers"
	"engineer/src/app/internal/middleware"
	"engineer/src/app/internal/pages"
	"engineer/src/app/internal/seo"
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

	r.Use(chimiddleware.Logger)
	r.Use(chimiddleware.Recoverer)
	r.Use(middleware.CookieMiddleware)

	fileServer := http.FileServer(http.Dir("./src/app/static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	for _, lang := range []string{"EN", "DE", "FR", "ES", "IT", "PT", "NL", "RU", "JA", "KO", "ZH", "AR"} {
		langCode := lang
		r.Get("/"+strings.ToLower(langCode), func(w http.ResponseWriter, r *http.Request) {
			handleHomeWithLang(w, r, langCode)
		})
	}

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		lang := middleware.GetLang(r)
		http.Redirect(w, r, "/"+strings.ToLower(lang), http.StatusTemporaryRedirect)
	})

	r.Post("/api/theme", handlers.SetTheme)
	r.Post("/api/lang", handlers.SetLanguage)

	r.Get("/robots.txt", seo.Robots)
	r.Get("/sitemap.xml", seo.Sitemap)
	r.Get("/llms.txt", seo.LLMs)

	port := getPort()
	fmt.Printf("Server starting on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func getPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}

	startPort := 3000
	for port := startPort; port < startPort+10; port++ {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			listener.Close()
			return addr
		}
	}

	return ":3000"
}

func handleHomeWithLang(w http.ResponseWriter, r *http.Request, langCode string) {
	theme := middleware.GetTheme(r)

	if r.URL.Query().Get("mode") == "engineer" {
		w.Header().Set("Content-Type", "text/html")
		_ = pages.EngineerPage(pages.HomePageProps{
			Theme:       theme,
			CurrentLang: langCode,
		}).Render(w)
		return
	}

	page := components.BaseLayout("4nuel", theme, langCode, seo.Tags(langCode),
		pages.HomePage(pages.HomePageProps{
			Theme:       theme,
			CurrentLang: langCode,
		}),
	)

	w.Header().Set("Content-Type", "text/html")
	_ = page.Render(w)
}
