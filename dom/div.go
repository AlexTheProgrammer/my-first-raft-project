package dom

type DivEl struct {
	*Element
}

func Div(innerHTML string, nodes ...Node) *DivEl {
	return &DivEl{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (d *DivEl) OpenTag() string  { return "<div>" }
func (d *DivEl) CloseTag() string { return "</div>" }
func (d *DivEl) IsNil() bool {
	return nil == d
}
