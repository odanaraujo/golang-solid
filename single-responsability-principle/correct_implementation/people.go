package correctimplementation

type People struct {
	Name string
	Age  int
}

func (p *People) GetPeopleId(id string) string {
	return "People"
}

func (p *People) AddPeople() *People {
	return &People{
		Name: p.Name,
		Age:  p.Age,
	}
}

func (p *People) UpdatePeople() *People {
	return &People{
		Name: p.Name,
		Age:  p.Age,
	}
}
