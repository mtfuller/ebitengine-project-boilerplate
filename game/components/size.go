package components

type Size struct {
	OffsetX, OffsetY, W, H int
}

func (s Size) GetName() string {
	return "size"
}
