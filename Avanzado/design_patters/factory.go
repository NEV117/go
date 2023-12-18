package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

// 电子书
type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getName() string {
	return c.name
}

func (c *Computer) getStock() int {
	return c.stock
}

type Labtop struct {
	Computer
}

func newLabtop() IProduct {
	return &Labtop{
		Computer: Computer{
			name:  "Laptop",
			stock: 10,
		},
	}
}

type Desktop struct {
	Computer
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

func GetComputerFactory(ComputerType string) (IProduct, error) {
	if ComputerType == "laptop" {
		return newLabtop(), nil
	}

	if ComputerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid Computer Type")
}

func (c *Computer) String() string {
	s := fmt.Sprintf("Product: %s, with Stock: %d", c.name, c.stock)
	return s
}

func PrintNameAndStock(p IProduct) {
	fmt.Printf("Product name: %s, with stock %d\n", p.getName(), p.getStock())
}

func mainn() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")

	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)
}
