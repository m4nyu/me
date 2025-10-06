package components

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// BaseLayout creates the base HTML document structure
func BaseLayout(title string, theme string, content ...g.Node) g.Node {
	// Default to system theme if not specified
	if theme == "" {
		theme = "system"
	}

	// Determine actual theme class
	themeClass := theme
	if theme == "system" {
		themeClass = "light" // Default for SSR
	}

	// Determine favicon path
	faviconPath := "/static/light.svg"
	if theme == "dark" {
		faviconPath = "/static/dark.svg"
	}

	bodyContent := append([]g.Node{Class("font-sans antialiased")}, content...)

	return Doctype(
		HTML(
			Lang("en"),
			Class(themeClass),
			g.Attr("suppressHydrationWarning", ""),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
				TitleEl(g.Text(title)),
				Link(Rel("icon"), Href(faviconPath)),

				// Tailwind CSS with dark mode config
				Script(Src("https://cdn.tailwindcss.com")),
				Script(g.Raw(`
					tailwind.config = {
						darkMode: 'class',
						theme: {
							extend: {
								colors: {
									background: 'var(--background)',
									foreground: 'var(--foreground)'
								}
							}
						}
					}
				`)),

				// Custom styles
				Link(Rel("stylesheet"), Href("/static/css/globals.css")),
			),
			Body(
				bodyContent...,
			),
		),
	)
}
