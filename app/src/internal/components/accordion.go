package components

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type AccordionItem struct {
	Question string
	Answer   string
}

// Accordion creates an accordion component with JavaScript-controlled exclusive opening
// Uses <details>/<summary> elements with a script to close others when one opens
func Accordion(class string, items []AccordionItem) g.Node {
	accordionItems := []g.Node{}

	for i, item := range items {
		accordionItems = append(accordionItems,
			AccordionItemComponent(i, item),
		)
	}

	return g.Group([]g.Node{
		// Minimal custom styles (only what Tailwind can't handle)
		g.El("style", g.Raw(`
			details summary::-webkit-details-marker {
				display: none;
			}

			.accordion-item .accordion-content {
				max-height: 0;
				transition: max-height 0.3s cubic-bezier(0.4, 0, 0.2, 1);
			}

			.accordion-item[open] summary .accordion-chevron {
				transform: rotate(180deg);
			}
		`)),
		Div(
			Class(class+" accordion-group"),
			g.Group(accordionItems),
		),
		// JavaScript to handle closing other accordions and smooth transitions
		g.El("script", g.Raw(`
			document.querySelectorAll('.accordion-group details').forEach(details => {
				const content = details.querySelector('.accordion-content');

				// Set initial height for smooth transitions
				if (details.open) {
					content.style.maxHeight = content.scrollHeight + 'px';
				}

				details.addEventListener('toggle', function() {
					const content = this.querySelector('.accordion-content');

					if (this.open) {
						// Close other accordions
						document.querySelectorAll('.accordion-group details').forEach(other => {
							if (other !== this && other.open) {
								const otherContent = other.querySelector('.accordion-content');
								otherContent.style.maxHeight = '0px';
								setTimeout(() => {
									other.open = false;
								}, 300);
							}
						});

						// Open this accordion with smooth transition
						content.style.maxHeight = '0px';
						setTimeout(() => {
							content.style.maxHeight = content.scrollHeight + 'px';
						}, 10);
					} else {
						// Close this accordion
						content.style.maxHeight = '0px';
					}
				});
			});
		`)),
	})
}

// AccordionItemComponent creates a single accordion item using details/summary
func AccordionItemComponent(index int, item AccordionItem) g.Node {
	isFirst := index == 0

	var detailsAttrs []g.Node
	detailsAttrs = append(detailsAttrs,
		Class("accordion-item overflow-hidden border-b border-foreground last:border-b-0"),
	)
	if isFirst {
		detailsAttrs = append(detailsAttrs, g.Attr("open"))
	}

	return g.El("details",
		g.Group(detailsAttrs),
		g.El("summary",
			Class("w-full text-left text-sm md:text-base font-light hover:text-foreground py-3 px-4 cursor-pointer text-foreground flex justify-between items-center list-none transition-all duration-200"),
			Span(g.Text(item.Question)),
			// Chevron icon
			g.El("svg",
				Class("w-4 h-4 accordion-chevron transition-transform duration-200"),
				g.Attr("fill", "none"),
				g.Attr("stroke", "currentColor"),
				g.Attr("viewBox", "0 0 24 24"),
				g.Attr("xmlns", "http://www.w3.org/2000/svg"),
				g.El("path",
					g.Attr("stroke-linecap", "round"),
					g.Attr("stroke-linejoin", "round"),
					g.Attr("stroke-width", "2"),
					g.Attr("d", "M19 9l-7 7-7-7"),
				),
			),
		),
		// Content
		Div(
			Class("accordion-content overflow-hidden text-foreground text-xs md:text-sm leading-relaxed font-light"),
			Div(
				Class("px-4 py-3"),
				P(g.Text(item.Answer)),
			),
		),
	)
}
