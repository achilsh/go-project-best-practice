package unit_test_demo

type Person struct {
	Age  int
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}
func (p *Person) SetName(n string) {
	p.Name = n
}
func (p *Person) SetOption(n string, a int) {
	p.Age = a
	p.Name = n
}
