package main

import d "mfrp/dom"

func main() {
	ch := make(chan bool)
	d := d.DOM{
		Body: d.Body(
			d.Div("",
				d.Div("",
					d.UL("",
						d.LI("",
							d.A("",
								d.SVG(""),
								d.Span("Dashboard"),
							),
						),
					),
				),
			),
		),
	}

	d.Render()

	<-ch // block indefinitely to keep program alive
}
