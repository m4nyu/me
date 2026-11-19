package components

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// BaseLayout renders the document shell. theme is "light", "dark" or "system";
// system sets no class, so CSS and the favicon follow the OS colour scheme.
func BaseLayout(title string, theme string, content ...g.Node) g.Node {
	themeClass := ""
	favicon := "/static/icon.svg" // adapts to the OS scheme by itself
	switch theme {
	case "light":
		themeClass, favicon = "light", "/static/light.svg"
	case "dark":
		themeClass, favicon = "dark", "/static/dark.svg"
	}

	bodyContent := append([]g.Node{Class("font-sans antialiased font-medium bg-background text-foreground select-none")}, content...)

	return Doctype(
		HTML(
			Lang("en"),
			Class(themeClass+" scroll-smooth"),
			g.Attr("style", "scroll-snap-type: y mandatory; scroll-behavior: smooth;"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
				Meta(Name("color-scheme"), Content("light dark")),
				TitleEl(g.Text(title)),
				Link(Rel("icon"), Type("image/svg+xml"), Href(favicon)),
				Link(Rel("stylesheet"), Href("/static/css/app.css")),
			),
			Body(bodyContent...),
		),
	)
}
