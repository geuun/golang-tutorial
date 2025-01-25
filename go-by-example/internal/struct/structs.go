package _struct

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 30

	// return pointer to the struct
	return &p
}

func Main() {
	fmt.Println(person{"geun", 30})
	fmt.Println(person{name: "geun", age: 30})
	fmt.Println(person{age: 30, name: "geun"})
	fmt.Println(person{name: "geun"})
	fmt.Println(&person{name: "geun", age: 30}) // pointer
	fmt.Println(newPerson("geun"))

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
