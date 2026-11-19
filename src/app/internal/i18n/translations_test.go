package i18n

import "testing"

func TestT(t *testing.T) {
	t.Run("returns english translation", func(t *testing.T) {
		got := T("EN", "tagline")
		if got == "" {
			t.Error("expected non-empty tagline for EN")
		}
	})

	t.Run("returns german translation", func(t *testing.T) {
		got := T("DE", "tagline")
		if got == "" {
			t.Error("expected non-empty tagline for DE")
		}
	})

	t.Run("unknown key returns key as fallback", func(t *testing.T) {
		got := T("EN", "nonexistent_key")
		if got != "nonexistent_key" {
			t.Errorf("expected key as fallback, got %q", got)
		}
	})

	t.Run("all languages have core keys", func(t *testing.T) {
		coreKeys := []string{"tagline", "intro_desc", "core_strengths", "opinions", "services_title", "web_title", "aiml_title"}
		for lang := range Translations {
			for _, key := range coreKeys {
				if T(lang, key) == "" {
					t.Errorf("language %q missing translation for %q", lang, key)
				}
			}
		}
	})
}
