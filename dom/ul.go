package dom

// Unordered List
type UL struct {
	*Element
}

func NewUL(innerHTML string, nodes ...Node) *UL {
	return &UL{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (u *UL) OpenTag() string  { return "<ul>" }
func (u *UL) CloseTag() string { return "</ul>" }
func (u *UL) IsNil() bool {
	return nil == u
}
