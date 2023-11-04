package dom

type BodyEl struct {
	Element
}

func Body(nodes ...Node) *BodyEl {
	return &BodyEl{
		Element{
			Ns: nodes,
		},
	}
}

func (b *BodyEl) Div(innerHTML string) *BodyEl {
	b.El(Div(innerHTML))

	return b
}

func (b *BodyEl) El(n Node) *BodyEl {
	b.Element.El(n)

	return b
}

func (b *BodyEl) OpenTag() string  { return "<body>" }
func (b *BodyEl) CloseTag() string { return "</body>" }
func (b *BodyEl) IsNil() bool {
	return nil == b
}
