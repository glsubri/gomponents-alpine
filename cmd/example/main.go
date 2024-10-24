package main

import (
	"net/http"

	x "github.com/glsubri/gomponents-alpine"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func main() {
	s := server{}
	http.ListenAndServe(":3000", &s)
}

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		displayPage().Render(w)
	default:
		http.Error(w, "Unsupported method.", http.StatusBadRequest)
	}
}

func displayPage() Node {
	return HtmlLayout(
		Div(
			Class("p-10 flex flex-col gap-10"),

			toggle(),
			maskInput(),
		),
	)
}

func HtmlLayout(children ...Node) Node {
	return Doctype(
		HTML(
			htmlLayoutHead(),
			Body(
				Class("overscroll-x-none"),

				Group(children),
			),
		),
	)
}

func htmlLayoutHead() Node {
	return Head(
		Meta(Charset("UTF-8")),
		Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),

		// Dependencies
		Script(Defer(), Src("https://cdnjs.cloudflare.com/ajax/libs/alpinejs-mask/3.14.3/cdn.js")),
		Script(Defer(), Src("https://cdn.jsdelivr.net/npm/alpinejs@3.14.3/dist/cdn.js")),

		Script(Defer(), Src("https://cdn.tailwindcss.com")),
	)
}

func toggle() Node {
	return Label(
		x.Data("{ isChecked: false }"),

		Class("relative inline-block h-8 w-14 cursor-pointer rounded-full transition [-webkit-tap-highlight-color:_transparent]"),
		x.Class("{ 'bg-gray-300': !isChecked, 'bg-green-500': isChecked }"),

		Input(
			x.Model("isChecked"),

			Type("checkbox"),
			Name("someFormInput"),
			Class("sr-only"),
		),

		Span(
			Class("absolute inset-y-0 m-1 inline-flex size-6 items-center justify-center rounded-full bg-white transition-all"),
			Attr(":class", "{ 'text-gray-400 start-0': !isChecked, 'start-6 text-green-600': isChecked }"),

			SVG(
				x.Show("!isChecked"),

				Attr("data-unchecked-icon"),
				Attr("xmlns", "http://www.w3.org/2000/svg"),
				Attr("viewBox", "0 0 20 20"),
				Attr("fill", "currentColor"),
				Class("size-4"),

				El("path",
					Attr("fill-rule", "evenodd"),
					Attr("d", "M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"),
					Attr("clip-rule", "evenodd"),
				),
			),

			SVG(
				x.Show("isChecked"),

				Attr("data-unchecked-icon"),
				Attr("xmlns", "http://www.w3.org/2000/svg"),
				Attr("viewBox", "0 0 20 20"),
				Attr("fill", "currentColor"),
				Class("size-4"),

				El("path",
					Attr("fill-rule", "evenodd"),
					Attr("d", "M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"),
					Attr("clip-rule", "evenodd"),
				),
			),
		),
	)
}

func maskInput() Node {
	return Input(
		x.Data(""),
		x.Mask("99/99/9999"),

		Type("text"),
		Placeholder("DD/MM/YYYY"),
		Class("w-96 border rounded p-2"),
	)
}
