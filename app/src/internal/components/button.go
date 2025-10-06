package components

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

type ButtonProps struct {
	Type      string
	Class     string
	OnClick   string
	Name      string
	Value     string
	AriaLabel string
}

// CustomButton creates a button element with customizable properties
func CustomButton(props ButtonProps, children ...g.Node) g.Node {
	attrs := []g.Node{}

	if props.Type != "" {
		attrs = append(attrs, Type(props.Type))
	} else {
		attrs = append(attrs, Type("button"))
	}

	if props.Class != "" {
		attrs = append(attrs, Class(props.Class))
	}

	if props.OnClick != "" {
		attrs = append(attrs, g.Attr("@click", props.OnClick))
	}

	if props.Name != "" {
		attrs = append(attrs, Name(props.Name))
	}

	if props.Value != "" {
		attrs = append(attrs, Value(props.Value))
	}

	if props.AriaLabel != "" {
		attrs = append(attrs, g.Attr("aria-label", props.AriaLabel))
	}

	attrs = append(attrs, children...)

	return Button(attrs...)
}
