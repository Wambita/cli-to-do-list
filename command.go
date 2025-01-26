package main

import "flag"

// commands
type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	list   bool
}

// initialze commnads
func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "Add", "", "Add a new to do specify title")
	flag.StringVar(&cf.Edit, "Edit", "", "Edit a to do by index and specify a new title id:new_title")
	flag.IntVar(&cf.Del, "Del", -1, "Specify a to do by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a to do by index to toggle completed status")
	flag.BoolVar(&cf.list, "list", false, "List all todos")

	//parse the arguments
	flag.Parse()

	//return ref to cmdFlag
	return &cf
}


//execute the flags
func (cf *CmdFlags)  Execute ( todos *Todos) {
	switch{
	case cf.list:
		todos.print()
	case cf.Add != "": 
		todos.add(cf.Add)
	case 		
	}
}