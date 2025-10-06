package handlers

import (
	"net/http"
)

// SetTheme handles setting the theme cookie via POST
func SetTheme(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	theme := r.FormValue("theme")
	if theme == "" {
		theme = "light"
	}

	// Validate theme value
	if theme != "light" && theme != "dark" && theme != "system" {
		theme = "system"
	}

	// Set cookie with 1 year expiration
	http.SetCookie(w, &http.Cookie{
		Name:     "theme",
		Value:    theme,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60, // 1 year
		HttpOnly: false,               // Allow JavaScript to read for client-side logic if needed
		SameSite: http.SameSiteLaxMode,
	})

	// Redirect back to referrer or home
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "/"
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}

// SetLanguage handles setting the language cookie via POST
func SetLanguage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lang := r.FormValue("lang")
	if lang == "" {
		lang = "EN"
	}

	// Validate language value
	validLangs := map[string]bool{
		"EN": true, "DE": true, "FR": true, "ES": true, "IT": true, "PT": true,
		"NL": true, "RU": true, "JA": true, "KO": true, "ZH": true, "AR": true,
	}
	if !validLangs[lang] {
		lang = "EN"
	}

	// Set cookie with 1 year expiration
	http.SetCookie(w, &http.Cookie{
		Name:     "lang",
		Value:    lang,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60, // 1 year
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	})

	// Redirect back to referrer or home
	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "/"
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}