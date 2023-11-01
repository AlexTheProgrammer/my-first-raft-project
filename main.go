package main

import "mfrp/dom"

func main() {
	ch := make(chan bool)
	d := dom.DOM{
		Body: dom.NewBody(
			dom.NewDiv("Hi there olly",
				dom.NewDiv("I made a nested div"),
			),
		),
	}

	d.Render()

	<-ch // block indefinitely to keep program alive
}
