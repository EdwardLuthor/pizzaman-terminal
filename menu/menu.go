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
			Manager.ListPizzas()

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
			pizzaManager = Manager.AddPizza(name, cost)
		case 3:
			fmt.Print("Введите ID пиццы для удаления: ")
			var id int
			var removed bool
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Println("Ошибка, некорректный формат ID")
				continue
			}
			pizzaManager, removed = Manager.RemovePizza(id)
			// found можно убрать, не имеет особого смысла
			if removed {
				fmt.Println("Пицца успешно удалена.")
			} else {
				fmt.Println("Пицца с таким ID не найдена. ")
			}
		case 0:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Некорректный выбор.")
		}
	}
}
