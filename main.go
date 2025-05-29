package main

import (
	"github.com/EdwardLuthor/pizzaman-terminal/menu"
	"github.com/EdwardLuthor/pizzaman-terminal/pizza"
)

func main() {
	pizzaManager := pizza.NewManager()
	menu.RunMenu(pizzaManager)
}
