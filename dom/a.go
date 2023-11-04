package dom

// a href
type A struct {
	*Element
}

func NewA(innerHTML string, nodes ...Node) *A {
	return &A{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (u *A) OpenTag() string  { return "<a>" }
func (u *A) CloseTag() string { return "</a>" }
func (u *A) IsNil() bool {
	return nil == u
}
