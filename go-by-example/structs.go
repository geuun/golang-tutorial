package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 30
	return &p
}

func main() {
	fmt.Println(person{"Geun", 31})

	s := person{name: "geun", age: 30}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"doggggggg",
		true,
	}
	fmt.Println(dog)
}
