package components

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// ModalDialog creates a modal dialog component
func ModalDialog(id string, title string, content ...g.Node) g.Node {
	return Div(
		g.Attr("x-data", "{ open: false }"),
		g.Attr("x-on:keydown.escape.window", "open = false"),
		// Dialog overlay
		Div(
			g.Attr("x-show", "open"),
			g.Attr("x-transition:enter", "transition ease-out duration-200"),
			g.Attr("x-transition:enter-start", "opacity-0"),
			g.Attr("x-transition:enter-end", "opacity-100"),
			g.Attr("x-transition:leave", "transition ease-in duration-150"),
			g.Attr("x-transition:leave-start", "opacity-100"),
			g.Attr("x-transition:leave-end", "opacity-0"),
			Class("fixed inset-0 z-50 bg-background/50"),
			g.Attr("x-on:click", "open = false"),
		),
		// Dialog content
		Div(
			g.Attr("x-show", "open"),
			g.Attr("x-transition:enter", "transition ease-out duration-200"),
			g.Attr("x-transition:enter-start", "opacity-0 scale-95"),
			g.Attr("x-transition:enter-end", "opacity-100 scale-100"),
			g.Attr("x-transition:leave", "transition ease-in duration-150"),
			g.Attr("x-transition:leave-start", "opacity-100 scale-100"),
			g.Attr("x-transition:leave-end", "opacity-0 scale-95"),
			Class("fixed top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 z-50 w-full max-w-lg bg-background border border-foreground rounded-none p-0 shadow-lg max-h-[80vh]"),
			Div(
				Class("overflow-y-auto max-h-[80vh] p-4 pt-2"),
				// Close button
				Button(
					Type("button"),
					g.Attr("x-on:click", "open = false"),
					Class("absolute top-4 right-4 z-50 cursor-pointer hover:opacity-70 transition-opacity p-1 h-auto rounded-none shadow-none border-0 bg-transparent hover:bg-transparent"),
					g.Attr("aria-label", "Close dialog"),
					// X icon (using simple SVG)
					g.El("svg",
						Class("w-4 h-4 text-foreground"),
						g.Attr("fill", "none"),
						g.Attr("stroke", "currentColor"),
						g.Attr("viewBox", "0 0 24 24"),
						g.El("path",
							g.Attr("stroke-linecap", "round"),
							g.Attr("stroke-linejoin", "round"),
							g.Attr("stroke-width", "2"),
							g.Attr("d", "M6 18L18 6M6 6l12 12"),
						),
					),
				),
				// Header
				Div(
					Class("flex flex-row items-center justify-between relative mt-2"),
					H2(Class("text-foreground font-light text-xl"), g.Text(title)),
				),
				// Content
				Div(
					Class("grid gap-6 py-4"),
					g.Group(content),
				),
			),
		),
	)
}
