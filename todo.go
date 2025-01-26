package main

import (
	"errors"
	"time"
)

// to do struct
type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

// to do list
type Todos []Todo

// func add that creates a to do and adds it to the list
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	// append a todo to the  todo list
	*todos = append(*todos, todo)
}

// validate if index provided for to functions is valid (helper function)
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		return errors.New("Invalid index")
	}
	return nil
}

// delete method
func (todos *Todos) delete(index int) error {
	
}

// func toggle that changes the status of todo completed T or F
// func (todos *Todos) toggle
