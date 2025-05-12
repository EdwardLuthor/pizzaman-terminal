package main

import (
	"fmt"

	"github.com/EdwardLuthor/pizzaman-terminal/blob/main/pizza"
)

func main() {
	pizza.AddPizza(pizza.Pizza{"Name": "Margarita", "Cost": 5.50})
	pizza.AddPizza(pizza.Pizza{"Name": "Pepperoni", "Cost": 6.75})
	pizza.AddPizza(pizza.Pizza{"Name": "Hawai", "Cost": 6.20})

	fmt.Println("Список пицц:")
	for _, p := range pizza.ListPizzas() {
		fmt.Printf("%d. %s - $%.2f\n", p["ID"], p["Name"], p["Cost"])
	}
	pizzaToRemoveID := 2
	if pizza.RemovePizza(pizzaToRemoveID) {
		fmt.Printf("\nПицца с ID %d удалена.\n", pizzaToRemoveID)
	} else {
		fmt.Printf("\nПицца с ID %d не найдена.\n", pizzaToRemoveID)
	}

	fmt.Println("\nСписок пицц после удаления:")
	for _, p := range pizza.ListPizzas() {
		fmt.Printf("%d. %s - $%.2f\n", p["ID"], p["Name"], p["Cost"])
	}
}
