package handlers

import (
	"net/http"
)

func SetTheme(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	theme := r.FormValue("theme")
	if theme == "" {
		theme = "light"
	}

	if theme != "light" && theme != "dark" && theme != "system" {
		theme = "system"
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "theme",
		Value:    theme,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60, // 1 year
		HttpOnly: false,               // Allow JavaScript to read for client-side logic if needed
		SameSite: http.SameSiteLaxMode,
	})

	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "/"
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}


func SetLanguage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	lang := r.FormValue("lang")
	if lang == "" {
		lang = "EN"
	}

	validLangs := map[string]bool{
		"EN": true, "DE": true, "FR": true, "ES": true, "IT": true, "PT": true,
		"NL": true, "RU": true, "JA": true, "KO": true, "ZH": true, "AR": true,
	}
	if !validLangs[lang] {
		lang = "EN"
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "lang",
		Value:    lang,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60, // 1 year
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	})

	referer := r.Header.Get("Referer")
	if referer == "" {
		referer = "/"
	}
	http.Redirect(w, r, referer, http.StatusSeeOther)
}