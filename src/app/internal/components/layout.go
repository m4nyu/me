package components

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func BaseLayout(title string, theme string, content ...g.Node) g.Node {
	if theme == "" {
		theme = "system"
	}

	themeClass := theme
	if theme == "system" {
		themeClass = "light" // Default for SSR
	}

	faviconPath := "/static/light.svg"
	if theme == "dark" {
		faviconPath = "/static/dark.svg"
	}

	bodyContent := append([]g.Node{Class("font-sans antialiased font-normal bg-background text-foreground select-none")}, content...)

	return Doctype(
		HTML(
			Lang("en"),
			Class(themeClass+" scroll-smooth"),
			g.Attr("suppressHydrationWarning", ""),
			g.Attr("style", "scroll-snap-type: y mandatory; scroll-behavior: smooth;"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
				TitleEl(g.Text(title)),
				Link(Rel("icon"), Href(faviconPath)),

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

				Link(Rel("stylesheet"), Href("/static/css/globals.css")),

				g.El("style", g.Raw(`
					/* Allow selection on interactive elements */
					button, a, label, input, select, textarea {
						user-select: auto;
						-webkit-user-select: auto;
						-moz-user-select: auto;
						-ms-user-select: auto;
						pointer-events: auto;
					}
				`)),
			),
			Body(
				append(bodyContent,
					Script(g.Raw(`
						(function() {
							let ws;
							let reloadOnReconnect = false;

							function connect() {
								ws = new WebSocket('ws://' + location.host + '/livereload');

								ws.onopen = function() {
									console.log('[LiveReload] Connected');
									if (reloadOnReconnect) {
										console.log('[LiveReload] Reloading page...');
										location.reload();
									}
								};

								ws.onclose = function() {
									console.log('[LiveReload] Disconnected. Reconnecting...');
									reloadOnReconnect = true;
									setTimeout(connect, 1000);
								};

								ws.onerror = function() {
									ws.close();
								};
							}

							connect();
						})();
					`)),
				)...,
			),
		),
	)
}
