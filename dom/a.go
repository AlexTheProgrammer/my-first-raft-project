package dom

// a href
type AEl struct {
	*Element
}

func NewA(innerHTML string, nodes ...Node) *AEl {
	return &AEl{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (u *AEl) OpenTag() string  { return "<a>" }
func (u *AEl) CloseTag() string { return "</a>" }
func (u *AEl) IsNil() bool {
	return nil == u
}
