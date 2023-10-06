package correctimplementation

import "fmt"

func (p *People) PrintPeople() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}
