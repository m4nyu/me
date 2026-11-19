package components

import (
	"strings"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Favicon returns the M icon for theme ("light", "dark" or "system"); the
// system one adapts to the OS colour scheme by itself.
func Favicon(theme string) string {
	switch theme {
	case "light":
		return "/static/light.svg"
	case "dark":
		return "/static/dark.svg"
	}
	return "/static/icon.svg"
}

// BaseLayout renders the document shell. theme is "light", "dark" or "system";
// system sets no class, so CSS and the favicon follow the OS colour scheme.
func BaseLayout(title, theme, lang string, head []g.Node, content ...g.Node) g.Node {
	themeClass := ""
	if theme == "light" || theme == "dark" {
		themeClass = theme
	}

	bodyContent := append([]g.Node{Class("font-sans antialiased font-medium bg-background text-foreground select-none")}, content...)

	return Doctype(
		HTML(
			Lang(strings.ToLower(lang)),
			Class(themeClass+" scroll-smooth"),
			g.Attr("style", "scroll-snap-type: y mandatory; scroll-behavior: smooth;"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
				Meta(Name("color-scheme"), Content("light dark")),
				TitleEl(g.Text(title)),
				Link(Rel("icon"), Type("image/svg+xml"), Href(Favicon(theme))),
				Link(Rel("stylesheet"), Href("/static/css/app.css")),
				g.Group(head),
			),
			Body(bodyContent...),
		),
	)
}
