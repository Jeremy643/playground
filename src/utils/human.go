package utils

type Human interface {
	GetName() string
	GetAge() int
}

type human struct {
	Name string
	Age  int
}

func (h *human) GetName() string {
	return h.Name
}

func (h *human) GetAge() int {
	return h.Age
}

func CreateHuman(name string, age int) *human {
	return &human{
		Name: name,
		Age:  age,
	}
}
