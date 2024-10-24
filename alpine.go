// Package alpine provides Alpine attributes and helpers for gomponents.
// See https://alpinejs.dev
package apline

import (
	"fmt"

	g "maragu.dev/gomponents"
)

// Data defines a chunk of HTML as an Alpine component and provides the reactive data for that component to reference.
// see https://alpinejs.dev/directives/data
func Data(v string) g.Node {
	return attr("data", v)
}

// Init allows you to hook into the initialization phase of any element in Alpine.
// see https://alpinejs.dev/directives/init
func Init(v string) g.Node {
	return attr("init", v)
}

// Show provides an expressive way to show and hide DOM elements.
// see https://alpinejs.dev/directives/show
func Show(v string) g.Node {
	return attr("show", v)
}

// Bind allows you to set HTML attributes on elements based on the result of JavaScript expressions.
// see https://alpinejs.dev/directives/bind
func Bind(onAttr, v string) g.Node {
	return attr(fmt.Sprintf("bind:%s", onAttr), v)
}

// On allows you to easily run code on dispatched DOM events.
// see https://alpinejs.dev/directives/on
func On(event, v string) g.Node {
	return attr(fmt.Sprintf("on:%s", event), v)
}

// Text sets the text content of an element to the result of a given expression.
// see https://alpinejs.dev/directives/text
func Text(v string) g.Node {
	return attr("text", v)
}

// Html sets the "innerHTML" property of an element to the result of a given expression.
// see https://alpinejs.dev/directives/html
func Html(v string) g.Node {
	return attr("html", v)
}

// Model allows you to bind the value of an input element to Alpine data.
// see https://alpinejs.dev/directives/model
func Model(v string) g.Node {
	return attr("model", v)
}

// Modelable allows you to expose any Alpine property as the target of the x-model directive.
// see https://alpinejs.dev/directives/modelable
func Modelable(v string) g.Node {
	return attr("modelable", v)
}

// For allows you to create DOM elements by iterating through a list.
// see https://alpinejs.dev/directives/for
func For(v string) g.Node {
	return attr("for", v)
}

// Transition allows you to create smooth transitions between when an element is shown or hidden.
// see https://alpinejs.dev/directives/transition
func Transition(v ...string) g.Node {
	var node g.Node
	switch len(v) {
	case 0:
		node = g.Attr("x-transition")
	case 1:
		node = g.Attr("x-transition" + v[0])
	default:
		mapped := make([]g.Node, len(v))
		for i, elem := range v {
			mapped[i] = g.Attr("x-transition" + elem)
		}

		return g.Group(mapped)
	}

	return node
}

// Effect is used for re-evaluating an expression when one of its dependencies change.
// see https://alpinejs.dev/directives/effect
func Effect(v string) g.Node {
	return attr("effect", v)
}

// Ignore will prevent Apline from parsing the element
// see https://alpinejs.dev/directives/ignore
func Ignore() g.Node {
	return attr("ignore")
}

// Ref in combination with $refs is a useful utility for easily accessing DOM elements directly
// see https://alpinejs.dev/directives/ref
func Ref(v string) g.Node {
	return attr("ref", v)
}

// Cloak hides the element it's attached to until Alpine is fully loaded on the page.
// see https://alpinejs.dev/directives/cloak
func Cloak() g.Node {
	return attr("cloak")
}

// Teleport allows you to transport part of your Alpine template to another part of the DOM on the page entirely.
// see https://alpinejs.dev/directives/teleport
func Teleport(v string) g.Node {
	return attr("teleport", v)
}

// If is used for toggling elements on the page, similarly to x-show, however it completely adds
// and removes the element it's applied to rather than just changing its CSS display property to "none".
// see https://alpinejs.dev/directives/if
func If(v string) g.Node {
	return attr("if", v)
}

// Id allows you to declare a new "scope" for any new IDs generated using $id().
// see https://alpinejs.dev/directives/id
func Id(v string) g.Node {
	return attr("id", v)
}

// Mask allows you to automatically format a text input field as a user types.
// Part of the Mask Plugin
// see https://alpinejs.dev/plugins/mask
func Mask(v string) g.Node {
	return attr("mask", v)
}

// MaskDynamic allows you to automatically format a text input field as a user types.
// Part of the Mask Plugin
// see https://alpinejs.dev/plugins/mask
func MaskDynamic(v string) g.Node {
	return attr("mask:dynamic", v)
}

func attr(name string, value ...string) g.Node {
	return g.Attr("x-"+name, value...)
}
