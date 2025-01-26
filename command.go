package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// commands
type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

// initialze commnads
func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "Add", "", "Add a new to do specify title")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a to do by index and specify a new title id:new_title")
	flag.IntVar(&cf.Del, "Del", -1, "Specify a to do by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a to do by index to toggle completed status")
	flag.BoolVar(&cf.list, "List", false, "List all todos")

	//parse the arguments
	flag.Parse()

	//return ref to cmdFlag
	return &cf
}


//execute the flags
func (cf *CmdFlags)  Execute ( todos *Todos) {
	switch{
	case cf.List:
		todos.print()
	case cf.Add != "": 
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error: invalid format for edit. Please use id:new_title")
			os.Exit(1)
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index,parts[1])
	case cf.Toggle != -1: 	
		todos.toggle(cf.Toggle)
	case cf.Del != -1:
		todos.delete(cf.Del)

	default:
		fmt.Println("Invalid command")	
	}


}