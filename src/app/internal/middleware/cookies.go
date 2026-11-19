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

func CookieMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		theme := "system"
		if themeCookie, err := r.Cookie("theme"); err == nil {
			if themeCookie.Value == "light" || themeCookie.Value == "dark" || themeCookie.Value == "system" {
				theme = themeCookie.Value
			}
		}

		lang := "EN"
		if langCookie, err := r.Cookie("lang"); err == nil {
			lang = langCookie.Value
		}

		ctx := context.WithValue(r.Context(), ThemeKey, theme)
		ctx = context.WithValue(ctx, LangKey, lang)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetTheme(r *http.Request) string {
	if theme, ok := r.Context().Value(ThemeKey).(string); ok {
		return theme
	}
	return "system"
}

func GetLang(r *http.Request) string {
	if lang, ok := r.Context().Value(LangKey).(string); ok {
		return lang
	}
	return "EN"
}

func detectLanguageFromHeader(r *http.Request) string {
	acceptLang := r.Header.Get("Accept-Language")
	if acceptLang == "" {
		return "EN"
	}

	languages := strings.Split(acceptLang, ",")
	for _, lang := range languages {
		lang = strings.Split(lang, ";")[0]
		lang = strings.TrimSpace(lang)
		primaryLang := strings.ToLower(strings.Split(lang, "-")[0])

		if supportedLang, ok := supportedLanguages[primaryLang]; ok {
			return supportedLang
		}
	}

	return "EN" // fallback
}
