package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCookieMiddleware(t *testing.T) {
	handler := CookieMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		theme := GetTheme(r)
		lang := GetLang(r)
		w.Header().Set("X-Theme", theme)
		w.Header().Set("X-Lang", lang)
	}))

	t.Run("defaults without cookies", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)

		if got := w.Header().Get("X-Theme"); got != "system" {
			t.Errorf("theme = %q, want %q", got, "system")
		}
		if got := w.Header().Get("X-Lang"); got != "EN" {
			t.Errorf("lang = %q, want %q", got, "EN")
		}
	})

	t.Run("reads theme cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "theme", Value: "dark"})
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)

		if got := w.Header().Get("X-Theme"); got != "dark" {
			t.Errorf("theme = %q, want %q", got, "dark")
		}
	})

	t.Run("reads lang cookie", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "lang", Value: "DE"})
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)

		if got := w.Header().Get("X-Lang"); got != "DE" {
			t.Errorf("lang = %q, want %q", got, "DE")
		}
	})

	t.Run("ignores invalid theme", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.AddCookie(&http.Cookie{Name: "theme", Value: "neon"})
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, req)

		if got := w.Header().Get("X-Theme"); got != "system" {
			t.Errorf("theme = %q, want %q", got, "system")
		}
	})
}

func TestGetThemeDefault(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if got := GetTheme(req); got != "system" {
		t.Errorf("GetTheme = %q, want %q", got, "system")
	}
}

func TestGetLangDefault(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	if got := GetLang(req); got != "EN" {
		t.Errorf("GetLang = %q, want %q", got, "EN")
	}
}

func TestDetectLanguageFromHeader(t *testing.T) {
	tests := []struct {
		name   string
		header string
		want   string
	}{
		{"empty", "", "EN"},
		{"english", "en-US,en;q=0.9", "EN"},
		{"german", "de-DE,de;q=0.9,en;q=0.8", "DE"},
		{"japanese", "ja,en;q=0.5", "JA"},
		{"unknown falls back", "xx-XX", "EN"},
		{"french first", "fr-FR,fr;q=0.9,en;q=0.8", "FR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			if tt.header != "" {
				req.Header.Set("Accept-Language", tt.header)
			}
			got := detectLanguageFromHeader(req)
			if got != tt.want {
				t.Errorf("detectLanguageFromHeader = %q, want %q", got, tt.want)
			}
		})
	}
}
