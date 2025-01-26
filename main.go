package main

func main() {
	//create todos list
	todos := Todos{}

	//store todos in  a json file
	storage := NewStorage[Todos]("todos.json")

	//populate table with json data
	storage.Load(&todos)

	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	//save the todolist
	storage.Save(todos)
}
