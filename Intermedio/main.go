package main

import (
	PK "Intermedio/mypackage"
	"fmt"
)

func main() {
	e2 := new(PK.Employee)
	e4 := PK.NewEmployee(2, "Nuevo empleado")
	//fmt.Printf("%v", e)
	e2.Id = 1
	e2.Name = "Name"
	//fmt.Printf("%v", e)
	e2.SetId(5)
	e2.SetName("Name 2")
	//fmt.Printf("%v", e)
	fmt.Println(e2.GetId())
	fmt.Println(e2.GetName())

	fmt.Println("e4.GetId()", e4.GetId())
	fmt.Println("e4.GetName()", e4.GetName())
}
