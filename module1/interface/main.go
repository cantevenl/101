package main

import "fmt"

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

func (h *Human) getName() string {
	return h.firstName + "," + h.lastName
}

func (p Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

type Car struct {
	factory, model string
}

func (c *Car) getName() string {
	return c.factory + "-" + c.model
}

type Animal struct {
	Name string
	Age  string
}

func (a *Animal) getName() string {
	return a.Name + "$" + a.Age
}

func main() {
	interfaces := []IF{}
	h := new(Human)

	var human1 IF
	human1 = new(Human)
	human1.getName()

	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}
	p := Plane{}
	p.vendor = "testVendor"
	p.model = "testModel"
	fmt.Println(p.getName())

	a := Animal{"pig", "3"}
	fmt.Println(a.getName())
}
