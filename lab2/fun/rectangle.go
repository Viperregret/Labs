package fun

type Rectange struct {
	Width  float32
	Height float32
}

func NewRectangle(width float32, height float32) Rectange {
	r := Rectange{Width: width, Height: height}

	return r
}

func Square(r Rectange) float32 {
	return r.Height * r.Width
}
