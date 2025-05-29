package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/EdwardLuthor/pizzaman-terminal/pizza"
)

type Menu struct {
	pizzaManager *pizza.Manager
	reader       *bufio.Reader
}

func NewMenu() *Menu {
	return &Menu{
		pizzaManager: pizza.NewManager(),
		reader:       bufio.NewReader(os.Stdin),
	}
}
func (m *Menu) displayMainMenu() {
	fmt.Println("Выберите действие:")
	fmt.Println("1. Вывести список пицц")
	fmt.Println("2. Добавить новую пиццу")
	fmt.Println("3. Удалить пиццу")
	fmt.Println("0. Выход")

	fmt.Print("> ")
}

func (m *Menu) getChoice() (int, error) {
	var choiceSTR string
	choiceSTR, err := m.reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("Ошибка чтения выбора: ", err)
	}
	choiceSTR = strings.TrimSpace(choiceSTR)

	choice, err := strconv.Atoi(choiceSTR)
	if err != nil {
		return 0, fmt.Errorf("Некорректный формат выбора: ", err)
	}
	return choice, nil
}
func (m *Menu) ListPizzas() {

	pizzas := m.pizzaManager.ListPizzas()
	fmt.Println("\n--- Меню пицц ---")
	if len(pizzas) == 0 {
		fmt.Println("В меню пока нет пицц.")
	} else {
		for _, p := range pizzas {
			fmt.Printf("%d. %s - $%.2f\n", p.ID, p.Name, p.Cost)
		}
	}
	fmt.Println("-------------------\n")
}

func (m *Menu) getPizzaDetailsFromUser() (name string, cost float64, err error) {
	fmt.Print("Введите название новой пиццы: ")
	name, err = m.reader.ReadString('\n')
	if err != nil {
		return "", 0, fmt.Errorf("ошибка чтения названия: ", err)
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return "", 0, fmt.Errorf("название пиццы не может быть пустым")
	}

	fmt.Print("Введите стоимость новой пиццы: ")
	costSTR, err := m.reader.ReadString('\n')
	if err != nil {
		return "", 0, fmt.Errorf("ошибка чтения стоимости: ", err)
	}
	costSTR = strings.TrimSpace(costSTR)

	cost, err = strconv.ParseFloat(costSTR, 64)
	if err != nil {
		return "", 0, fmt.Errorf("некорректный формат стоимости. Пожалуйста, введите число.")
	}
	return name, cost, nil
}

func (m *Menu) AddpPizza() {
	name, cost, err := m.getPizzaDetailsFromUser()
	if err != nil {
		fmt.Println("Ошибка добавления пиццы:", err)
		return
	}

	pizzaID := m.pizzaManager.AddPizza(name, cost)
	fmt.Printf("Пицца '%s' добавлена в меню с ID %d.\n\n", name, pizzaID)
}

func (m *Menu) getPizzaIDFromUser() (id int, err error) {
	fmt.Print("Введите ID пиццы для удаления: ")
	idSTR, err := m.reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения ID: ", err)
	}
	idSTR = strings.TrimSpace(idSTR)

	id, err = strconv.Atoi(idSTR)
	if err != nil {
		return 0, fmt.Errorf("некорректный формат ID. Пожалуйста, введите число.")
	}
	return id, nil
}

func (m *Menu) RemovePizza() {
	id, err := m.getPizzaIDFromUser()
	if err != nil {
		fmt.Println("Ошибка удаления пиццы:", err)
		return
	}

	wasRemoved := m.pizzaManager.RemovePizza(id)

	if wasRemoved {
		fmt.Printf("Пицца с ID %d успешно удалена.\n\n", id)
	} else {
		fmt.Printf("Пицца с ID %d не найдена в меню.\n\n", id)
	}
}

func RunMenu() {

	menu := NewMenu()

	for {
		menu.displayMainMenu()
		choice, err := menu.getChoice()
		if err != nil {
			fmt.Println("Ошибка", err)
			continue
		}

		switch choice {
		case 1:
			menu.ListPizzas()
		case 2:
			menu.AddpPizza()
		case 3:
			menu.RemovePizza()
		case 0:
			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Некорректный выбор.")
		}
	}
}
