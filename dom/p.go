package dom

// Support Vector Graphic
type P struct {
	*Element
}

func NewP(innerHTML string, nodes ...Node) *P {
	return &P{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (b *P) OpenTag() string  { return "<svg>" }
func (b *P) CloseTag() string { return "</svg>" }
func (b *P) IsNil() bool {
	return nil == b
}
