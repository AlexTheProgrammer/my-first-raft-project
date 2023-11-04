package dom

type Span struct {
	*Element
}

func NewSpan(innerHTML string, nodes ...Node) *Span {
	return &Span{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (b *Span) OpenTag() string  { return "<span>" }
func (b *Span) CloseTag() string { return "</span>" }
func (b *Span) IsNil() bool {
	return nil == b
}
