package middleware

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	ThemeKey contextKey = "theme"
	LangKey  contextKey = "lang"
)

// Supported languages mapping from ISO 639-1 codes to our language codes
var supportedLanguages = map[string]string{
	"en": "EN",
	"de": "DE",
	"fr": "FR",
	"es": "ES",
	"it": "IT",
	"pt": "PT",
	"nl": "NL",
	"ru": "RU",
	"ja": "JA",
	"ko": "KO",
	"zh": "ZH",
	"ar": "AR",
}

// CookieMiddleware reads theme and language from cookies with fallbacks
func CookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Read theme from cookie, default to system
		theme := "system"
		if themeCookie, err := r.Cookie("theme"); err == nil {
			if themeCookie.Value == "light" || themeCookie.Value == "dark" || themeCookie.Value == "system" {
				theme = themeCookie.Value
			}
		}

		// Language is detected from cookie first, then Accept-Language header
		lang := "EN"
		if langCookie, err := r.Cookie("lang"); err == nil {
			lang = langCookie.Value
		} else {
			lang = detectLanguageFromHeader(r)
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
	return "system"
}

// GetLang retrieves the language from the request context
func GetLang(r *http.Request) string {
	if lang, ok := r.Context().Value(LangKey).(string); ok {
		return lang
	}
	return "EN"
}

// detectLanguageFromHeader parses the Accept-Language header and returns the best matching language
func detectLanguageFromHeader(r *http.Request) string {
	acceptLang := r.Header.Get("Accept-Language")
	if acceptLang == "" {
		return "EN"
	}

	// Parse Accept-Language header (format: "en-US,en;q=0.9,de;q=0.8")
	languages := strings.Split(acceptLang, ",")
	for _, lang := range languages {
		// Remove quality factor (;q=0.9)
		lang = strings.Split(lang, ";")[0]
		// Remove whitespace
		lang = strings.TrimSpace(lang)
		// Extract primary language code (en-US -> en)
		primaryLang := strings.ToLower(strings.Split(lang, "-")[0])

		// Check if we support this language
		if supportedLang, ok := supportedLanguages[primaryLang]; ok {
			return supportedLang
		}
	}

	return "EN" // fallback
}
