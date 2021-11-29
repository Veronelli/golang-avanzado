package main

import "fmt"

type Topic interface {
	ragistrar(observer Observer)
	broadcast()
}

type Observer interface {
	getId() string
	updateValue(string)
}

type Item struct {
	observers []Observer
	name      string
	avaible   bool
}

func NewItem(name string) *Item {
	return &Item{
		name: name,
	}
}

func (i *Item) register(observer Observer) {
	i.observers = append(i.observers, observer)
}

func (i *Item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}

}

func (i *Item) UpdateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.avaible = true
	i.broadcast()
}

type EmailClient struct {
	id string
}

func (e *EmailClient) updateValue(value string) {
	fmt.Printf("Sending Email - %s available from client %s \n", value, e.id)
}

func (e *EmailClient) getId() string {
	return e.getId()

}

func main() {
	nvideaItem := NewItem("RTX 3080")
	firstObserver := &EmailClient{
		id: "12ab",
	}

	secondObserver := &EmailClient{
		id: "34dc",
	}

	nvideaItem.register(firstObserver)
	nvideaItem.register(secondObserver)
	nvideaItem.UpdateAvailable()
}
