package dom

// Unordered List
type ULEl struct {
	*Element
}

func NewUL(innerHTML string, nodes ...Node) *ULEl {
	return &ULEl{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (u *ULEl) OpenTag() string  { return "<ul>" }
func (u *ULEl) CloseTag() string { return "</ul>" }
func (u *ULEl) IsNil() bool {
	return nil == u
}
