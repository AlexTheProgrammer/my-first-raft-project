package main

import (
	"github.com/AlexTheProgrammer/sail/dom"
)

func main() {
	ch := make(chan bool)
	d := dom.DOM{
		Body: dom.NewBody(
			dom.NewDiv("Hi there olly"),
		),
	}

	d.Render()

	<-ch // block indefinitely to keep program alive
}
