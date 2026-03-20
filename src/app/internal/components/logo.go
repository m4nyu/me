package components

import (
	"fmt"

	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Logo(size string) g.Node {
	sizeClass := "w-16 h-16 sm:w-20 sm:h-20 md:w-24 md:h-24 lg:w-28 lg:h-28 xl:w-32 xl:h-32"
	if size == "small" {
		sizeClass = "w-8 h-8"
	} else if size == "large" {
		sizeClass = "w-20 h-20 sm:w-24 sm:h-24 md:w-28 md:h-28 lg:w-32 lg:h-32 xl:w-36 xl:h-36"
	} else if size == "xlarge" {
		sizeClass = "w-24 h-24 sm:w-28 sm:h-28 md:w-32 md:h-32 lg:w-36 lg:h-36 xl:w-40 xl:h-40"
	}

	gridSize := 6
	tileBaseSize := 100.0 / float64(gridSize)

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
					g.Attr("viewBox", "0 0 256 256"),
					g.Attr("style", fmt.Sprintf(
						"position: absolute; width: %.2f%%; height: %.2f%%; left: -%.2f%%; top: -%.2f%%;",
						(100.0/tileBaseSize)*100, (100.0/tileBaseSize)*100, (x/tileBaseSize)*100, (y/tileBaseSize)*100,
					)),
					g.El("rect",
						g.Attr("width", "256"),
						g.Attr("height", "256"),
						Class("fill-foreground"),
					),
					g.El("g",
						g.Attr("transform", "translate(128,128) rotate(40.5) translate(-128,-128)"),
						g.El("path",
							g.Attr("d", "M68 192 L68 64 L84 64 L121.6 144 L134.4 144 L172 64 L188 64 L188 192 L172 192 L172 84 L137.6 164 L118.4 164 L84 84 L84 192 Z"),
							Class("fill-background"),
						),
					),
				),
			))
		}
	}

	// Clean single SVG shown after animation completes
	cleanSVG := g.El("svg",
		g.Attr("width", "100%"),
		g.Attr("height", "100%"),
		g.Attr("viewBox", "0 0 256 256"),
		g.Attr("class", "logo-clean absolute inset-0 w-full h-full"),
		g.El("rect",
			g.Attr("width", "256"),
			g.Attr("height", "256"),
			Class("fill-foreground"),
		),
		g.El("g",
			g.Attr("transform", "translate(128,128) rotate(40.5) translate(-128,-128)"),
			g.El("path",
				g.Attr("d", "M68 192 L68 64 L84 64 L121.6 144 L134.4 144 L172 64 L188 64 L188 192 L172 192 L172 84 L137.6 164 L118.4 164 L84 84 L84 192 Z"),
				Class("fill-background"),
			),
		),
	)

	return g.Group([]g.Node{
		g.El("style", g.Raw(`
			@keyframes tile-reveal {
				from { opacity: 0; }
				to { opacity: 1; }
			}

			.logo-tile {
				animation: tile-reveal 0.1s ease-in forwards;
				opacity: 0;
			}

			.logo-clean {
				opacity: 0;
				transition: opacity 0.15s ease-in;
			}

			.logo-clean.visible {
				opacity: 1;
			}

			.animated-logo .logo-tile:nth-child(1) { animation-delay: 0.826s; }
			.animated-logo .logo-tile:nth-child(2) { animation-delay: 0.028s; }
			.animated-logo .logo-tile:nth-child(3) { animation-delay: 0.704s; }
			.animated-logo .logo-tile:nth-child(4) { animation-delay: 0.441s; }
			.animated-logo .logo-tile:nth-child(5) { animation-delay: 0.530s; }
			.animated-logo .logo-tile:nth-child(6) { animation-delay: 0.409s; }
			.animated-logo .logo-tile:nth-child(7) { animation-delay: 1.441s; }
			.animated-logo .logo-tile:nth-child(8) { animation-delay: 0.028s; }
			.animated-logo .logo-tile:nth-child(9) { animation-delay: 0.053s; }
			.animated-logo .logo-tile:nth-child(10) { animation-delay: 0.658s; }
			.animated-logo .logo-tile:nth-child(11) { animation-delay: 1.355s; }
			.animated-logo .logo-tile:nth-child(12) { animation-delay: 0.471s; }
			.animated-logo .logo-tile:nth-child(13) { animation-delay: 0.926s; }
			.animated-logo .logo-tile:nth-child(14) { animation-delay: 1.343s; }
			.animated-logo .logo-tile:nth-child(15) { animation-delay: 0.495s; }
			.animated-logo .logo-tile:nth-child(16) { animation-delay: 1.078s; }
			.animated-logo .logo-tile:nth-child(17) { animation-delay: 0.827s; }
			.animated-logo .logo-tile:nth-child(18) { animation-delay: 0.959s; }
			.animated-logo .logo-tile:nth-child(19) { animation-delay: 1.948s; }
			.animated-logo .logo-tile:nth-child(20) { animation-delay: 0.640s; }
			.animated-logo .logo-tile:nth-child(21) { animation-delay: 0.696s; }
			.animated-logo .logo-tile:nth-child(22) { animation-delay: 0.368s; }
			.animated-logo .logo-tile:nth-child(23) { animation-delay: 1.615s; }
			.animated-logo .logo-tile:nth-child(24) { animation-delay: 0.539s; }
			.animated-logo .logo-tile:nth-child(25) { animation-delay: 1.173s; }
			.animated-logo .logo-tile:nth-child(26) { animation-delay: 0.585s; }
			.animated-logo .logo-tile:nth-child(27) { animation-delay: 1.323s; }
			.animated-logo .logo-tile:nth-child(28) { animation-delay: 1.196s; }
			.animated-logo .logo-tile:nth-child(29) { animation-delay: 0.384s; }
			.animated-logo .logo-tile:nth-child(30) { animation-delay: 1.879s; }
			.animated-logo .logo-tile:nth-child(31) { animation-delay: 1.269s; }
			.animated-logo .logo-tile:nth-child(32) { animation-delay: 1.844s; }
			.animated-logo .logo-tile:nth-child(33) { animation-delay: 1.042s; }
			.animated-logo .logo-tile:nth-child(34) { animation-delay: 0.227s; }
			.animated-logo .logo-tile:nth-child(35) { animation-delay: 1.801s; }
			.animated-logo .logo-tile:nth-child(36) { animation-delay: 1.645s; }
		`)),
		Div(
			Class(sizeClass+" flex-shrink-0 relative animated-logo"),
			cleanSVG,
			g.Group(tiles),
		),
		g.El("script", g.Raw(`
			(function() {
				setTimeout(function() {
					document.querySelectorAll('.animated-logo').forEach(function(logo) {
						var clean = logo.querySelector('.logo-clean');
						if (clean) clean.classList.add('visible');
						logo.querySelectorAll('.logo-tile').forEach(function(t) { t.remove(); });
					});
				}, 2200);
			})();
		`)),
	})
}
