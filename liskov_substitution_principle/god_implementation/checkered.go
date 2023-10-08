package godimplementation

type Area interface {
	GetArea() float64
}

type RectangleType interface {
	GetWidth() float64
	GetHeight() float64
	Area
}

type SquareType interface {
	Area
	SetSize(size float64)
	GetSize() float64
}

type Square struct {
	Size float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) GetArea() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) GetWidth() float64 {
	return r.Width
}

func (r *Rectangle) GetHeight() float64 {
	return r.Height
}

func (r *Rectangle) SetWidth(width float64) {
	r.Width = width
}

func (r *Rectangle) SetHeight(height float64) {
	r.Height = height
}

func (s *Square) GetArea() float64 {
	return s.Size * s.Size
}

func (s *Square) GetSize() float64 {
	return s.Size
}

func (s *Square) SetSize(size float64) {
	s.Size = size
}
