package components

import (
	"engineer/src/app/internal/i18n"
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

func ThemeSwitcher(currentTheme string) g.Node {
	themeLabel := currentTheme
	if currentTheme == "system" {
		themeLabel = "auto"
	}

	return Div(
		Class("relative inline-flex mode-dropdown-container"),
		Button(
			Type("button"),
			Class("text-foreground text-xs sm:text-sm font-light cursor-pointer transition-all flex items-center gap-1 sm:gap-2 p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] mode-dropdown-trigger"),
			g.Attr("aria-label", "Select mode"),
			g.Attr("onclick", "this.parentElement.querySelector('.mode-dropdown-menu').classList.toggle('hidden')"),
			Span(
				Class("text-[10px] sm:text-xs font-light text-foreground capitalize"),
				g.Text(themeLabel),
			),
			g.Raw(`<svg xmlns="http://www.w3.org/2000/svg" width="10" height="10" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" class="text-foreground"><path d="m6 9 6 6 6-6"/></svg>`),
		),
		Div(
			Class("mode-dropdown-menu hidden absolute bottom-full left-0 mb-1 bg-background border border-foreground z-50 min-w-[100px]"),
			modeOption("light", currentTheme),
			modeOption("dark", currentTheme),
			modeOption("auto", currentTheme),
			A(
				Href("?mode=engineer"),
				Class("block w-full text-left px-3 py-1.5 text-[10px] sm:text-xs font-light text-foreground hover:bg-foreground hover:text-background transition-colors cursor-pointer"),
				g.Text("engineer"),
			),
		),
		g.El("script", g.Raw(`
			document.addEventListener('click', function(e) {
				document.querySelectorAll('.mode-dropdown-menu').forEach(function(m) {
					if (!m.parentElement.contains(e.target)) m.classList.add('hidden');
				});
			});
		`)),
	)
}

func modeOption(mode, currentTheme string) g.Node {
	themeValue := mode
	if mode == "auto" {
		themeValue = "system"
	}
	activeClass := ""
	if (mode == "auto" && currentTheme == "system") || mode == currentTheme {
		activeClass = " bg-foreground text-background"
	}
	return FormEl(
		Method("post"),
		Action("/api/theme"),
		Input(Type("hidden"), Name("theme"), Value(themeValue)),
		Button(
			Type("submit"),
			Class("block w-full text-left px-3 py-1.5 text-[10px] sm:text-xs font-light text-foreground hover:bg-foreground hover:text-background transition-colors cursor-pointer"+activeClass),
			g.Text(mode),
		),
	)
}

func LanguageSwitcher(currentLang string) g.Node {
	return g.El("label",
		g.Attr("for", "language-drawer"),
		Class("text-foreground hover:text-foreground text-xs sm:text-sm font-light cursor-pointer transition-all flex items-center gap-1 sm:gap-2 p-0 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent hover:[text-shadow:_0_0_10px_rgba(0,0,0,0.3)] dark:hover:[text-shadow:_0_0_10px_rgba(255,255,255,0.5)] active:[text-shadow:_0_0_15px_rgba(0,0,0,0.5)] dark:active:[text-shadow:_0_0_15px_rgba(255,255,255,0.7)]"),
		Span(
			Class("inline-flex rounded-md px-1 sm:px-2 py-0 text-[10px] sm:text-xs font-light border-0 text-foreground"),
			g.Text(currentLang),
		),
	)
}

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
			Input(
				Type("checkbox"),
				ID("language-drawer"),
				Class("drawer-toggle"),
				g.Attr("style", "display: none;"),
			),
		Div(
			Class("drawer-overlay hidden fixed inset-0 bg-black/50 z-50 opacity-0 transition-opacity duration-300"),
			ID("language-drawer-overlay"),
			g.Attr("aria-label", "Close drawer"),
		),
		Div(
			Class("drawer-content fixed bottom-0 left-0 right-0 z-[51] bg-background border-t border-foreground max-h-[80vh] translate-y-full transition-transform duration-300 ease-[cubic-bezier(0.32,0.72,0,1)]"),
			ID("language-drawer-content"),
			Div(
				ID("drawer-drag-handle"),
				Class("pb-2"),
				Div(
					Class("mx-auto mt-4 mb-2 h-1 w-[100px] rounded-full bg-foreground"),
				),
				Div(
					Class("flex flex-row items-center justify-between relative px-4"),
					H2(
						Class("text-foreground font-light text-lg"),
						g.Text(i18n.T(currentLang, "choose_language")),
					),
				),
			),
			Div(
				Class("overflow-y-auto max-h-[calc(80vh-8rem)] p-4 pt-2"),
				Div(
					Class("grid grid-cols-4 gap-1"),
					g.Group(languageItems),
				),
			),
		),
		g.El("script", g.Raw(`
			(function() {
				const drawer = document.getElementById('language-drawer-content');
				const handle = document.getElementById('drawer-drag-handle');
				const toggle = document.getElementById('language-drawer');
				const overlay = document.getElementById('language-drawer-overlay');
				if (!drawer || !handle || !toggle || !overlay) return;

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
