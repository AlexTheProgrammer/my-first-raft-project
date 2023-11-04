package dom

// Support Vector Graphic
type SVG struct {
	*Element
}

func NewSVG(innerHTML string, nodes ...Node) *SVG {
	return &SVG{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (b *SVG) OpenTag() string  { return "<svg>" }
func (b *SVG) CloseTag() string { return "</svg>" }
func (b *SVG) IsNil() bool {
	return nil == b
}
