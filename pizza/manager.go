package pizza

type Manager struct {
	pizzas []Pizza
	nextID int
}

func NewManager() *Manager {
	return &Manager{pizzas: make([]Pizza, 0),
		nextID: 1}
}

func (m Manager) ListPizzas() []Pizza {
	return m.pizzas
}

func (m *Manager) AddPizza(name string, cost float64) int {
	newPizza := Pizza{
		ID:   m.nextID,
		Name: name,
		Cost: cost,
	}
	m.pizzas = append(m.pizzas, newPizza)
	m.nextID++
	return newPizza.ID
}

func (m *Manager) RemovePizza(id int) bool {
	newPizzas := []Pizza{}
	var found bool = false

	for _, p := range m.pizzas {
		if p.ID != id {
			newPizzas = append(newPizzas, p)
		} else {
			found = true
		}
	}
	m.pizzas = newPizzas
	return found
}
