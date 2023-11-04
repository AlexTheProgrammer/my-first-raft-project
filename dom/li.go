package dom

// Unordered List
type LIEl struct {
	*Element
}

func NewLI(innerHTML string, nodes ...Node) *LIEl {
	return &LIEl{
		Element: NewElement(innerHTML, nodes...),
	}
}

func (u *LIEl) OpenTag() string  { return "<li>" }
func (u *LIEl) CloseTag() string { return "</li>" }
func (u *LIEl) IsNil() bool {
	return nil == u
}
