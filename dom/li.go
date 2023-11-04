package dom

// Unordered List
type LI struct {
	*Element
}

func NewLI(innerHTML string, nodes ...Node) *LI {
	return &LI{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (u *LI) OpenTag() string  { return "<li>" }
func (u *LI) CloseTag() string { return "</li>" }
func (u *LI) IsNil() bool {
	return nil == u
}
