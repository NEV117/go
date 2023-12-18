package mypackage

type Employee struct {
	Id   int
	Name string
}

func (e *Employee) SetId(id int) {
	e.Id = id
}

func (e *Employee) SetName(name string) {
	e.Name = name
}

func (e *Employee) GetId() int {
	return e.Id
}

func (e *Employee) GetName() string {
	return e.Name
}

// Constructor
func NewEmployee(id int, name string) *Employee {
	return &Employee{
		Id:   id,
		Name: name,
	}
}
