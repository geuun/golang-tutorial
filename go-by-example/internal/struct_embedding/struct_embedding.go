package _struct_embedding

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

type person struct {
	Name string
	Age  int
}

type employee struct {
	person  // struct embedding
	Company string
	Salary  float64
}

func Main() {
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	fmt.Println("also num:", co.base.num)

	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("describer:", d.describe())

	emp := employee{
		person: person{
			Name: "geun",
			Age:  29,
		},
		Company: "geuuuuuun inc",
		Salary:  9999999999999999,
	}

	// Access fields from person struct
	fmt.Println("Name:", emp.Name) // person.Name directly accessible
	fmt.Println("Age:", emp.Age)   // person.Age directly accessible

	fmt.Println(emp.person.Name)
	fmt.Println(emp.person.Age)
}
