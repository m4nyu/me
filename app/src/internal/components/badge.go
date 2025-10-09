package components

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Badge creates a badge component
func Badge(variant string, classes string, content ...g.Node) g.Node {
	baseClass := "inline-flex items-center rounded-md px-2 py-0 text-xs font-light"

	variantClass := ""
	switch variant {
	case "outline":
		variantClass = "text-foreground border-0"
	case "secondary":
		variantClass = "border-transparent bg-secondary text-secondary-foreground"
	case "destructive":
		variantClass = "border-transparent bg-destructive text-destructive-foreground"
	default:
		variantClass = "border-transparent bg-primary text-primary-foreground"
	}

	return Div(
		Class(baseClass+" "+variantClass+" "+classes),
		g.Group(content),
	)
}
