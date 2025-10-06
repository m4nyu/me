package components

import (
	"fmt"
	"math/rand"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Logo creates the animated logo component with tile reveal effect
func Logo(size string) g.Node {
	sizeClass := "w-16 h-16 sm:w-20 sm:h-20 md:w-24 md:h-24 lg:w-28 lg:h-28 xl:w-32 xl:h-32"
	if size == "small" {
		sizeClass = "w-8 h-8"
	} else if size == "large" {
		sizeClass = "w-20 h-20 sm:w-24 sm:h-24 md:w-28 md:h-28 lg:w-32 lg:h-32 xl:w-36 xl:h-36"
	} else if size == "xlarge" {
		sizeClass = "w-24 h-24 sm:w-28 sm:h-28 md:w-32 md:h-32 lg:w-36 lg:h-36 xl:w-40 xl:h-40"
	}

	// Generate tiles with random delays (matching Next.js seeded random)
	gridSize := 6
	tileBaseSize := 100.0 / float64(gridSize)

	// Seeded random to match Next.js
	seed := int64(12345)
	seededRandom := func() float64 {
		seed = (seed*9301 + 49297) % 233280
		return float64(seed) / 233280.0
	}

	tiles := []g.Node{}
	for row := 0; row < gridSize; row++ {
		for col := 0; col < gridSize; col++ {
			delay := seededRandom() * 2000
			x := float64(col) * tileBaseSize
			y := float64(row) * tileBaseSize

			tiles = append(tiles, Div(
				Class("absolute logo-tile"),
				g.Attr("data-delay", fmt.Sprintf("%.0f", delay)),
				g.Attr("style", fmt.Sprintf(
					"width: calc(%.2f%% + 1px); height: calc(%.2f%% + 1px); left: %.2f%%; top: %.2f%%; overflow: hidden;",
					tileBaseSize, tileBaseSize, x, y,
				)),
				g.El("svg",
					g.Attr("width", "100%"),
					g.Attr("height", "100%"),
					g.Attr("viewBox", "0 0 32 32"),
					g.Attr("shape-rendering", "crispEdges"),
					g.Attr("style", fmt.Sprintf(
						"position: absolute; width: %.2f%%; height: %.2f%%; left: -%.2f%%; top: -%.2f%%; image-rendering: -webkit-optimize-contrast; image-rendering: crisp-edges;",
						(100.0/tileBaseSize)*100, (100.0/tileBaseSize)*100, (x/tileBaseSize)*100, (y/tileBaseSize)*100,
					)),
					g.El("rect",
						g.Attr("width", "32"),
						g.Attr("height", "32"),
						Class("fill-foreground"),
					),
					g.El("g",
						g.Attr("transform", "translate(16,16) rotate(40.5) translate(-16,-16)"),
						g.El("path",
							g.Attr("d", "M8.5 24 L8.5 8 L10.5 8 L15.2 18 L16.8 18 L21.5 8 L23.5 8 L23.5 24 L21.5 24 L21.5 10.5 L17.2 20.5 L14.8 20.5 L10.5 10.5 L10.5 24 L8.5 24 Z"),
							Class("fill-background"),
							g.Attr("shape-rendering", "crispEdges"),
						),
					),
				),
			))
		}
	}

	return Div(
		Class(sizeClass+" flex-shrink-0 relative animated-logo"),
		g.Group(tiles),
	)
}

func init() {
	rand.Seed(12345)
}
