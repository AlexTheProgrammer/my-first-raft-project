package main

import (
	"github.com/AlexTheProgrammer/sail/dom"
)

func main() {
	ch := make(chan bool)
	d := dom.DOM{
		Body: []dom.Node{
			&dom.Div{InnerHTML: "Hi there sail and raft"},
		},
	}
	d.Render()

	<-ch // block indefinitely to keep program alive
}
