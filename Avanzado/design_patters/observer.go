package main

import "fmt"

type Topic interface {
	register(ibserver Observer)
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	available bool
}

func NewItem(name string) *Item {
	return &Item{name: name}
}
func (i *Item) updateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

func (i *Item) register(Observer Observer) {
	i.observers = append(i.observers, Observer)
}

type EmailCLient struct {
	id string
}

func (eC *EmailCLient) updateValue(value string) {
	fmt.Printf("Sending Email -%s available from client %s\n", value, eC.id)
}

func (eC *EmailCLient) getId() string {
	return eC.id
}

func main() {
	nvidiaItem := NewItem("RTX 3080")
	firstObserver := &EmailCLient{
		id: "12ab",
	}
	secondObserver := &EmailCLient{
		id: "34dc",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.updateAvailable()
}
