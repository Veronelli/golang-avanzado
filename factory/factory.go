package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) setName(name string) {
	c.name = name
}
func (c *Computer) getName() string {
	return c.name
}

type Laptop struct {
	Computer
}

func NewLaptop() IProduct {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 2000,
		},
	}

}

type Desktop struct {
	Computer
}

func NewDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Dell i5",
			stock: 1000,
		},
	}

}

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return NewLaptop(), nil
	} else if computerType == "desktop" {
		return NewDesktop(), nil
	}
	return nil, fmt.Errorf("invalid computer type")

}

func PrintNameAndStock(product IProduct) {
	fmt.Printf("Name: %s, Stock: %d \n", product.getName(), product.getStock())

}

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")
	PrintNameAndStock(laptop)
	PrintNameAndStock(desktop)

}
