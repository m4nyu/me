package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSetTheme(t *testing.T) {
	tests := []struct {
		name     string
		theme    string
		wantCook string
	}{
		{"light", "light", "light"},
		{"dark", "dark", "dark"},
		{"system", "system", "system"},
		{"invalid defaults to system", "invalid", "system"},
		{"empty defaults to light", "", "light"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"theme": {tt.theme}}
			req := httptest.NewRequest(http.MethodPost, "/api/theme", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Referer", "/en")
			w := httptest.NewRecorder()

			SetTheme(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusSeeOther {
				t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusSeeOther)
			}

			cookies := resp.Cookies()
			if len(cookies) == 0 {
				t.Fatal("expected theme cookie")
			}
			if cookies[0].Value != tt.wantCook {
				t.Errorf("got cookie %q, want %q", cookies[0].Value, tt.wantCook)
			}
		})
	}
}

func TestSetThemeRejectsGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/theme", nil)
	w := httptest.NewRecorder()
	SetTheme(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("got status %d, want %d", w.Code, http.StatusMethodNotAllowed)
	}
}

func TestSetLanguage(t *testing.T) {
	tests := []struct {
		name     string
		lang     string
		wantCook string
	}{
		{"valid EN", "EN", "EN"},
		{"valid DE", "DE", "DE"},
		{"valid JA", "JA", "JA"},
		{"invalid defaults to EN", "XX", "EN"},
		{"empty defaults to EN", "", "EN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{"lang": {tt.lang}}
			req := httptest.NewRequest(http.MethodPost, "/api/lang", strings.NewReader(form.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("Referer", "/en")
			w := httptest.NewRecorder()

			SetLanguage(w, req)

			resp := w.Result()
			if resp.StatusCode != http.StatusSeeOther {
				t.Errorf("got status %d, want %d", resp.StatusCode, http.StatusSeeOther)
			}

			cookies := resp.Cookies()
			if len(cookies) == 0 {
				t.Fatal("expected lang cookie")
			}
			if cookies[0].Value != tt.wantCook {
				t.Errorf("got cookie %q, want %q", cookies[0].Value, tt.wantCook)
			}
		})
	}
}

func TestSetLanguageRejectsGet(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/lang", nil)
	w := httptest.NewRecorder()
	SetLanguage(w, req)
	if w.Code != http.StatusMethodNotAllowed {
		t.Errorf("got status %d, want %d", w.Code, http.StatusMethodNotAllowed)
	}
}

func TestSetThemeRedirectsToRoot(t *testing.T) {
	form := url.Values{"theme": {"dark"}}
	req := httptest.NewRequest(http.MethodPost, "/api/theme", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()

	SetTheme(w, req)

	loc := w.Result().Header.Get("Location")
	if loc != "/" {
		t.Errorf("got redirect %q, want %q", loc, "/")
	}
}
