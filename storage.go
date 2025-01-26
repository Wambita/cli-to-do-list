package main

import (
	"encoding/json"
	"os"
)

// generics
// allow you to write code that can work with any data type while maintaining type safety. They allow you to write functions, types, and data structures that can operate on different types without needing to specify exactly what types will be used at the time of writing the code.
// A generic function that accepts any type and returns the same value
// func Print[T any](value T) {
//     fmt.Println(value)
// }

// func main() {
//     Print(42)      // Prints: 42
//     Print("Hello") // Prints: Hello
//     Print(3.14)    // Prints: 3.14
// }

type Storage[T any] struct {
	FileName string
}

// constructor  special type of function or method used to initialize an object or data structure.
// create a constructor is by using a function that returns a pointer to a newly created object or data structure.
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

//save
func(s *Storage[T]) save(data T) error {
	//convert data into json
	fileData, err := json.MarshalIndent(data, "","    ")
	if err != nil {
		return err
	}
	//write data 
	return os.WriteFile(s.FileName, fileData, 0644)
}
//load