package components

import (
	"engineer/app/src/internal/i18n"
	"strings"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

var languages = []struct {
	Code string
	Name string
}{
	{"EN", "English"},
	{"DE", "Deutsch"},
	{"FR", "Français"},
	{"ES", "Español"},
	{"IT", "Italiano"},
	{"PT", "Português"},
	{"NL", "Nederlands"},
	{"RU", "Русский"},
	{"JA", "日本語"},
	{"KO", "한국어"},
	{"ZH", "中文"},
	{"AR", "العربية"},
}

// ThemeSwitcher creates a theme mode switcher (server-side with cookies)
func ThemeSwitcher(currentTheme string) g.Node {
	// Determine next theme in cycle: light -> dark -> system -> light
	nextTheme := "light"
	if currentTheme == "light" {
		nextTheme = "dark"
	} else if currentTheme == "dark" {
		nextTheme = "system"
	}

	// Theme icon based on current theme
	var themeIcon g.Node
	themeLabel := currentTheme
	if currentTheme == "system" {
		themeLabel = "Auto"
	}

	if currentTheme == "light" {
		themeIcon = g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="sm:w-4 sm:h-4 text-foreground"><circle cx="12" cy="12" r="4"/><path d="M12 2v2"/><path d="M12 20v2"/><path d="m4.93 4.93 1.41 1.41"/><path d="m17.66 17.66 1.41 1.41"/><path d="M2 12h2"/><path d="M20 12h2"/><path d="m6.34 17.66-1.41 1.41"/><path d="m19.07 4.93-1.41 1.41"/></svg>`)
	} else if currentTheme == "dark" {
		themeIcon = g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="sm:w-4 sm:h-4 text-foreground"><path d="M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z"/></svg>`)
	} else {
		themeIcon = g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" class="sm:w-4 sm:h-4 text-foreground"><path d="M12 8V4H8"/><rect width="16" height="12" x="4" y="8" rx="2"/><path d="M2 14h2"/><path d="M20 14h2"/><path d="M15 13v2"/><path d="M9 13v2"/></svg>`)
	}

	return FormEl(
		Method("post"),
		Action("/api/theme"),
		Class("inline-flex"),
		Input(Type("hidden"), Name("theme"), Value(nextTheme)),
		Button(
			Type("submit"),
			Class("text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-all flex items-center gap-1 sm:gap-2 p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"),
			g.Attr("aria-label", "Toggle dark/light/auto mode"),
			Div(
				Class("w-[14px] h-[14px] sm:w-4 sm:h-4 flex items-center justify-center"),
				themeIcon,
			),
			Span(
				Class("inline-flex rounded-md px-1 sm:px-2 py-0 text-[10px] sm:text-xs font-light border-0 text-foreground capitalize"),
				g.Text(themeLabel),
			),
		),
	)
}

// LanguageSwitcher creates just the trigger for the language drawer
func LanguageSwitcher(currentLang string) g.Node {
	return g.El("label",
		g.Attr("for", "language-drawer"),
		Class("text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-all flex items-center gap-1 sm:gap-2 p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"),
		LanguagesIcon(),
		Span(
			Class("inline-flex rounded-md px-1 sm:px-2 py-0 text-[10px] sm:text-xs font-light border-0 text-foreground"),
			g.Text(currentLang),
		),
	)
}

// LanguageDrawer creates the drawer for language selection
func LanguageDrawer(currentLang string) g.Node {
	languageItems := []g.Node{}

	for _, lang := range languages {
		isSelected := currentLang == lang.Code
		bgClass := "bg-transparent border-foreground"
		textClass := "text-foreground group-hover:text-background"

		if isSelected {
			bgClass = "bg-foreground border-foreground"
			textClass = "text-background"
		}

		languageItems = append(languageItems,
			A(
				Href("/"+strings.ToLower(lang.Code)),
				Class("group relative border p-2 hover:bg-foreground transition-all duration-200 cursor-pointer flex items-center gap-2 h-auto rounded-none shadow-none "+bgClass),
				Div(
					Class("text-xs font-bold min-w-[1.25rem] text-center transition-colors "+textClass),
					g.Text(lang.Code),
				),
				Div(
					Class("text-[9px] font-medium transition-colors whitespace-nowrap "+textClass),
					g.Text(lang.Name),
				),
			),
		)
	}

	return g.Group([]g.Node{
		// Minimal custom styles (only checkbox state selectors that Tailwind can't handle)
		g.El("style", g.Raw(`
			.drawer-toggle:checked ~ .drawer-overlay {
				display: block;
				opacity: 1;
			}

			.drawer-toggle:checked ~ .drawer-content {
				transform: translateY(0);
			}
		`)),
		Div(
			Class("relative"),
			// Hidden checkbox to control drawer state
			Input(
				Type("checkbox"),
				ID("language-drawer"),
				Class("drawer-toggle"),
				g.Attr("style", "display: none;"),
			),
		// Overlay (clicking closes drawer)
		Div(
			Class("drawer-overlay hidden fixed inset-0 bg-black/50 z-50 opacity-0 transition-opacity duration-300"),
			ID("language-drawer-overlay"),
			g.Attr("aria-label", "Close drawer"),
		),
		// Drawer content
		Div(
			Class("drawer-content fixed bottom-0 left-0 right-0 z-[51] bg-background border-t border-foreground max-h-[80vh] translate-y-full transition-transform duration-300 ease-[cubic-bezier(0.32,0.72,0,1)]"),
			ID("language-drawer-content"),
			// Draggable header area
			Div(
				ID("drawer-drag-handle"),
				Class("pb-2"),
				// Handle bar
				Div(
					Class("mx-auto mt-4 mb-2 h-1 w-[100px] rounded-full bg-foreground"),
				),
				// Title
				Div(
					Class("flex flex-row items-center justify-between relative px-4"),
					H2(
						Class("text-foreground font-light text-lg"),
						g.Text(i18n.T(currentLang, "choose_language")),
					),
				),
			),
			// Drawer body
			Div(
				Class("overflow-y-auto max-h-[calc(80vh-8rem)] p-4 pt-2"),
				Div(
					Class("grid grid-cols-4 gap-1"),
					g.Group(languageItems),
				),
			),
		),
		// JavaScript for drag-to-close and overlay click
		g.El("script", g.Raw(`
			(function() {
				const drawer = document.getElementById('language-drawer-content');
				const handle = document.getElementById('drawer-drag-handle');
				const toggle = document.getElementById('language-drawer');
				const overlay = document.getElementById('language-drawer-overlay');
				if (!drawer || !handle || !toggle || !overlay) return;

				// Overlay click to close
				overlay.addEventListener('click', function() {
					toggle.checked = false;
				});

				let startY = 0;
				let currentY = 0;
				let startTime = 0;
				let isDragging = false;

				handle.addEventListener('touchstart', function(e) {
					startY = e.touches[0].clientY;
					currentY = startY;
					startTime = Date.now();
					isDragging = true;
				});

				handle.addEventListener('touchmove', function(e) {
					if (!isDragging) return;
					currentY = e.touches[0].clientY;
					const diffY = currentY - startY;

					if (diffY > 0) {
						e.preventDefault();
						drawer.style.transform = 'translateY(' + diffY + 'px)';
						drawer.style.transition = 'none';
					}
				});

				handle.addEventListener('touchend', function(e) {
					if (!isDragging) return;
					isDragging = false;

					const endTime = Date.now();
					const diffY = currentY - startY;
					const timeDiff = endTime - startTime;
					const velocity = Math.abs(diffY) / timeDiff;

					// Close if dragged down >80px OR fast swipe (velocity > 0.4)
					if (diffY > 80 || (velocity > 0.4 && diffY > 30)) {
						toggle.checked = false;
					}

					drawer.style.transition = '';
					drawer.style.transform = '';
				});
			})();
		`)),
		),
	})
}

// Icon components
func SunIcon() g.Node {
	return g.El("svg",
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Attr("width", "14"),
		g.Attr("height", "14"),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
		Class("sm:w-4 sm:h-4 text-foreground"),
		g.El("circle", g.Attr("cx", "12"), g.Attr("cy", "12"), g.Attr("r", "4")),
		g.El("path", g.Attr("d", "M12 2v2")),
		g.El("path", g.Attr("d", "M12 20v2")),
		g.El("path", g.Attr("d", "m4.93 4.93 1.41 1.41")),
		g.El("path", g.Attr("d", "m17.66 17.66 1.41 1.41")),
		g.El("path", g.Attr("d", "M2 12h2")),
		g.El("path", g.Attr("d", "M20 12h2")),
		g.El("path", g.Attr("d", "m6.34 17.66-1.41 1.41")),
		g.El("path", g.Attr("d", "m19.07 4.93-1.41 1.41")),
	)
}

func MoonIcon() g.Node {
	return g.El("svg",
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Attr("width", "14"),
		g.Attr("height", "14"),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
		Class("sm:w-4 sm:h-4 text-foreground"),
		g.El("path", g.Attr("d", "M12 3a6 6 0 0 0 9 9 9 9 0 1 1-9-9Z")),
	)
}

func LanguagesIcon() g.Node {
	return g.El("svg",
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Attr("width", "14"),
		g.Attr("height", "14"),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
		Class("sm:w-4 sm:h-4 text-foreground"),
		g.El("path", g.Attr("d", "m5 8 6 6")),
		g.El("path", g.Attr("d", "m4 14 6-6 2-3")),
		g.El("path", g.Attr("d", "M2 5h12")),
		g.El("path", g.Attr("d", "M7 2h1")),
		g.El("path", g.Attr("d", "m22 22-5-10-5 10")),
		g.El("path", g.Attr("d", "M14 18h6")),
	)
}

func XIcon() g.Node {
	return g.El("svg",
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Attr("width", "16"),
		g.Attr("height", "16"),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
		Class("w-4 h-4 text-foreground"),
		g.El("path", g.Attr("d", "M18 6 6 18")),
		g.El("path", g.Attr("d", "m6 6 12 12")),
	)
}

func BotIcon() g.Node {
	return g.El("svg",
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Attr("width", "14"),
		g.Attr("height", "14"),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
		Class("sm:w-4 sm:h-4 text-foreground"),
		g.El("path", g.Attr("d", "M12 8V4H8")),
		g.El("rect", g.Attr("width", "16"), g.Attr("height", "12"), g.Attr("x", "4"), g.Attr("y", "8"), g.Attr("rx", "2")),
		g.El("path", g.Attr("d", "M2 14h2")),
		g.El("path", g.Attr("d", "M20 14h2")),
		g.El("path", g.Attr("d", "M15 13v2")),
		g.El("path", g.Attr("d", "M9 13v2")),
	)
}

func ArrowLeftIcon() g.Node {
	return g.El("svg",
		g.Attr("xmlns", "http://www.w3.org/2000/svg"),
		g.Attr("width", "20"),
		g.Attr("height", "20"),
		g.Attr("viewBox", "0 0 24 24"),
		g.Attr("fill", "none"),
		g.Attr("stroke", "currentColor"),
		g.Attr("stroke-width", "2"),
		g.Attr("stroke-linecap", "round"),
		g.Attr("stroke-linejoin", "round"),
		Class("text-foreground"),
		g.El("path", g.Attr("d", "m12 19-7-7 7-7")),
		g.El("path", g.Attr("d", "M19 12H5")),
	)
}
