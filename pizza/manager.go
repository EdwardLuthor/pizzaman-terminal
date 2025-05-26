package pizza

import "fmt"

type Manager struct {
	pizzas []Pizza
}

func NewManager() Manager {
	return Manager{pizzas: make([]Pizza, 0)}
}

func (m Manager) ListPizzas() {
	fmt.Println("\n Меню пицц ")
	if len(m.pizzas) == 0 {
		fmt.Println("В меню пока нет пицц.")
		return
	}
	for _, p := range m.pizzas {
		fmt.Printf("%d. %s - $%.2f\n", p.ID, p.Name, p.Cost)
	}
}

func (m Manager) AddPizza(name string, cost float64) Manager {
	newPizza := Pizza{
		ID:   len(m.pizzas) + 1,
		Name: name,
		Cost: cost,
	}
	newManager := Manager{pizzas: append(m.pizzas, newPizza)}
	fmt.Printf("Пицца %s добавлена в меню с ID %d.\n\n", newPizza.Name, newPizza.ID)
	return newManager
}

func (m Manager) RemovePizza(id int) (Manager, bool) {
	initialLen := len(m.pizzas)
	newPizzas := []Pizza{}
	var found bool = false
	for _, p := range m.pizzas {
		if p.ID != id {
			newPizzas = append(newPizzas, p)
		} else {
			found = true
		}
	}
	newManager := Manager{pizzas: newPizzas}
	if len(newManager.pizzas) < initialLen {
		fmt.Printf("Пицца с ID %d удалена.\n\n", id)
	} else {
		fmt.Printf("Пицца с ID %d не найдена в меню.\n\n", id)
	}
	return newManager, found
}
