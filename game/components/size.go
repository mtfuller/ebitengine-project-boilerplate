package components

type Size struct {
	W, H int
}

func (s Size) GetName() string {
	return "size"
}