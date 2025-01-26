package main

import "time"

//to do struct
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

//to do list
type Todos []Todo

//func add that creates a to do and adds it to the list
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	//append a todo to the  todo list
	*todos = append(*todos, todo)
}
