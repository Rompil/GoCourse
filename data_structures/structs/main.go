package main

import "fmt"

//I create my own type
type Person struct { //this is an example of incapsulation in Go
	firstName string
	lastName  string
	age       int
}
type Stuff struct { //stuff inherits a method from person
	Person
	fired bool
}

//I could make
type число int

//and it would be an alias for int
func (p Person) fullInfo() string {
	return fmt.Sprintf("Name: %v \nSurname: %v\nAge: %v \n", p.firstName, p.lastName, p.age)
}

func main() {
	p := Person{"John", "Smith", 41}
	mrs := &Person{"Helen", "Miller", 18} //seems like we create a pointer
	boss := Stuff{
		Person: Person{
			firstName: "Jeff",
			lastName:  "Wallbert",
			age:       37,
		},
		fired: true,
	}
	fmt.Printf(p.fullInfo())
	fmt.Printf(mrs.fullInfo()) //but it still acts like a value variable. Go resolve pointer to value!!!
	fmt.Printf(boss.fullInfo())

	//	var n число =2
	//	m:= 2
	//	fmt.Println(n+m)//even if I made the alias for int I still cannot operate число and int in one operation
}
