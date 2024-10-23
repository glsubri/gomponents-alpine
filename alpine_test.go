package apline_test

import (
	"fmt"
	"testing"

	x "github.com/glsubri/gomponents-alpine"
	"github.com/glsubri/gomponents-alpine/internal/assert"

	g "maragu.dev/gomponents"
)

func TestAttributesWithString(t *testing.T) {
	cases := map[string]func(string) g.Node{
		// Alpine directives
		"data":      x.Data,
		"init":      x.Init,
		"show":      x.Show,
		"text":      x.Text,
		"html":      x.Html,
		"model":     x.Model,
		"modelable": x.Modelable,
		"for":       x.For,
		"effect":    x.Effect,
		"ref":       x.Ref,
		"teleport":  x.Teleport,
		"if":        x.If,
		"id":        x.Id,
		// Mask Plugin directives
		"mask":         x.Mask,
		"mask:dynamic": x.MaskDynamic,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output x-%v="balloon"`, name), func(t *testing.T) {
			n := g.El("div", fn("balloon"))
			assert.Equal(t, fmt.Sprintf(`<div x-%v="balloon"></div>`, name), n)
		})
	}
}

func TestAttributesWithNoArgs(t *testing.T) {
	cases := map[string]func() g.Node{
		"cloak":      x.Cloak,
		"transition": x.Transition,
		"ignore":     x.Ignore,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output x-%v"`, name), func(t *testing.T) {
			n := g.El("div", fn())
			assert.Equal(t, fmt.Sprintf(`<div x-%v></div>`, name), n)
		})
	}
}

func TestAttributesWithAttrAndValue(t *testing.T) {
	cases := map[string]func(string, string) g.Node{
		"bind": x.Bind,
		"on":   x.On,
	}

	for name, fn := range cases {
		t.Run(fmt.Sprintf(`should output x-%v:hot-air"`, name), func(t *testing.T) {
			n := g.El("div", fn("hot-air", "balloon"))
			assert.Equal(t, fmt.Sprintf(`<div x-%v:hot-air="balloon"></div>`, name), n)
		})
	}
}
