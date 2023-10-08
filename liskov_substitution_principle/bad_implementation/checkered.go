package badimplementation

type Checkered interface {
	GetArea() float64

	GetWidth() float64
	GetHeight() float64

	SetWidth(width float64)
	SetHeight(height float64)
}

type Square struct {
	Width  float64
	Height float64
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
	// devemos usar s.Height ou s.Width aqui?
	return s.Height * s.Height
}

func (s *Square) GetWidth() float64 {
	return s.Width
}

func (s *Square) GetHeight() float64 {
	return s.Height
}

func (s *Square) SetWidth(width float64) {
	//os dois lados do quadrado devem ser iguais
	s.Width = width
	s.Height = width
}

func (s *Square) SetHeight(height float64) {
	//ou devemos fazer r.SetWidth(height) aqui?
	s.Height = height
	s.Width = height
}
