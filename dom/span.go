package dom

type SpanEl struct {
	*Element
}

func Span(innerHTML string, nodes ...Node) *SpanEl {
	return &SpanEl{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (b *SpanEl) OpenTag() string  { return "<span>" }
func (b *SpanEl) CloseTag() string { return "</span>" }
func (b *SpanEl) IsNil() bool {
	return nil == b
}
