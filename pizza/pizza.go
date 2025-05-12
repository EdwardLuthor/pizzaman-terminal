package pizza

type Pizza map[string]interface{}

var pizzas []Pizza

func ListPizzas() []Pizza {
	return pizzas
}

func AddPizza(pizza Pizza) {
	pizza["ID"] = len(pizzas) + 1
	pizzas = append(pizzas, pizza)
}

func RemovePizza(id int) bool {
	newPizzas := []Pizza{}
	found := false
	for _, p := range pizzas {
		if p["ID"] != id {
			newPizzas = append(newPizzas, p)
		} else {
			found = true
		}
	}
	pizzas = newPizzas
	return found
}
