package dom

// Support Vector Graphic
type SVGEl struct {
	*Element
}

func SVG(innerHTML string, nodes ...Node) *SVGEl {
	return &SVGEl{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (b *SVGEl) OpenTag() string  { return "<svg>" }
func (b *SVGEl) CloseTag() string { return "</svg>" }
func (b *SVGEl) IsNil() bool {
	return nil == b
}
