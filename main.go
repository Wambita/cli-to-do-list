package main

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy bread")

	todos.delete(0)
	todos.toggle(0)
	todos.print()
}
