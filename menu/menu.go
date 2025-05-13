package menu

import (
	"fmt"

	"github.com/EdwardLuthor/pizzaman-terminal/pizza"
)

type PizzaManager struct {
	pizzas []pizza.Pizza
}

func NewPizzaManager() PizzaManager {
	return PizzaManager{pizzas: []pizza.Pizza{}}
}

func (m PizzaManager) ListPizzas() {
	fmt.Println("\n Меню пицц ")
	if len(m.pizzas) == 0 {
		fmt.Println("В меню пока нет пицц.")
		return
	}
	for _, p := range m.pizzas {
		fmt.Printf("%d. %s - $%.2f\n", p.ID, p.Name, p.Cost)
	}
}

func (m PizzaManager) AddPizza(name string, cost float64) PizzaManager {
	newPizza := pizza.Pizza{
		ID:   len(m.pizzas) + 1,
		Name: name,
		Cost: cost,
	}
	newManager := PizzaManager{pizzas: append(m.pizzas, newPizza)}
	fmt.Printf("Пицца %s добавлена в меню с ID %d.\n\n", newPizza.Name, newPizza.ID)
	return newManager
}

func (m PizzaManager) RemovePizza(id int) (PizzaManager, bool) {
	initialLen := len(m.pizzas)
	newPizzas := []pizza.Pizza{}
	var found bool = false
	for _, p := range m.pizzas {
		if p.ID != id {
			newPizzas = append(newPizzas, p)
		} else {
			found = true
		}
	}
	newManager := PizzaManager{pizzas: newPizzas}
	if len(newManager.pizzas) < initialLen {
		fmt.Printf("Пицца с ID %d удалена.\n\n", id)
	} else {
		fmt.Printf("Пицца с ID %d не найдена в меню.\n\n", id)
	}
	return newManager, found
}

func RunMenu() {
	pizzaManager := NewPizzaManager()
	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1. Вывести список пицц")
		fmt.Println("2. Добавить новую пиццу")
		fmt.Println("3. Удалить пиццу")
		fmt.Println("0. Выход")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода")
			continue
		}

		switch choice {
		case 1:
			pizzaManager.ListPizzas()

		case 2:
			fmt.Println("Введите название новой пиццы: ")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Введите стоимость новой пиццы: ")
			var cost float64
			_, err := fmt.Scanln(&cost)
			if err != nil {
				fmt.Println("Ошибка, некорректный формат стоимости")
				continue
			}
			pizzaManager = pizzaManager.AddPizza(name, cost)
		case 3:
			fmt.Print("Введите ID пиццы для удаления: ")
			var id int
			var removed bool
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println("Ошибка, некорректный формат ID")
				continue
			}
			pizzaManager, removed = pizzaManager.RemovePizza(id)
			if removed {
				fmt.Println("Пицца успешно удалена.")
			} else {
				fmt.Println("Пицца с таким ID не найдена. ")
			}
		case 0:
			fmt.Println("Выход из программы.")
		}
	}
}
