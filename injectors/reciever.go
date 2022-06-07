package injectors

import "fmt"

type Human struct {
	Name string
}

func injectors() {

	human := Human{
		Name: "John",
	}
	human.saludar()

}

func (h *Human) saludar() {
	fmt.Println("hi", h.Name)
}
