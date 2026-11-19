package components

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Logo renders the M mark as one vector SVG: crisp at every size and zoom.
// It fades in with CSS only (no script); reduced-motion users get it at once.
func Logo(size string) g.Node {
	sizeClass := "w-16 h-16 sm:w-20 sm:h-20 md:w-24 md:h-24 lg:w-28 lg:h-28 xl:w-32 xl:h-32"
	switch size {
	case "small":
		sizeClass = "w-8 h-8"
	case "large":
		sizeClass = "w-20 h-20 sm:w-24 sm:h-24 md:w-28 md:h-28 lg:w-32 lg:h-32 xl:w-36 xl:h-36"
	case "xlarge":
		sizeClass = "w-24 h-24 sm:w-28 sm:h-28 md:w-32 md:h-32 lg:w-36 lg:h-36 xl:w-40 xl:h-40"
	}

	return g.El("svg",
		g.Attr("viewBox", "0 0 256 256"),
		g.Attr("role", "img"),
		g.Attr("aria-label", "M"),
		g.Attr("class", sizeClass+" shrink-0 logo-mark"),
		g.El("rect", g.Attr("width", "256"), g.Attr("height", "256"), Class("fill-foreground")),
		g.El("path",
			g.Attr("transform", "rotate(40.5 128 128)"),
			g.Attr("d", "M68 192 L68 64 L84 64 L121.6 144 L134.4 144 L172 64 L188 64 L188 192 L172 192 L172 84 L137.6 164 L118.4 164 L84 84 L84 192 Z"),
			Class("fill-background"),
		),
	)
}
