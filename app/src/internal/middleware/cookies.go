package middleware

import (
	"context"
	"net/http"
)

type contextKey string

const (
	ThemeKey contextKey = "theme"
	LangKey  contextKey = "lang"
)

// CookieMiddleware reads theme and language cookies and adds them to the request context
func CookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read theme cookie
		theme := "light" // default
		if cookie, err := r.Cookie("theme"); err == nil {
			theme = cookie.Value
		}

		// Read language cookie
		lang := "EN" // default
		if cookie, err := r.Cookie("lang"); err == nil {
			lang = cookie.Value
		}

		// Add to context
		ctx := context.WithValue(r.Context(), ThemeKey, theme)
		ctx = context.WithValue(ctx, LangKey, lang)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetTheme retrieves the theme from the request context
func GetTheme(r *http.Request) string {
	if theme, ok := r.Context().Value(ThemeKey).(string); ok {
		return theme
	}
	return "light"
}

// GetLang retrieves the language from the request context
func GetLang(r *http.Request) string {
	if lang, ok := r.Context().Value(LangKey).(string); ok {
		return lang
	}
	return "EN"
}