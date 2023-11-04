package main

import "mfrp/dom"

func main() {
	ch := make(chan bool)
	d := dom.DOM{
		Body: dom.NewBody(
			dom.NewDiv("",
				dom.NewDiv("",
					dom.NewUL("",
						dom.NewLI("",
							dom.NewA("",
								dom.NewSVG(""),
								dom.NewSpan("Dashboard"),
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
