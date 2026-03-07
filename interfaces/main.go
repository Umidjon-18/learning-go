package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct {
	Name  string
	Age   int
	Color string
}

func (d Dog) String() string {
	return fmt.Sprintf("Dog name: %s, age: %d, color: %s", d.Name, d.Age, d.Color)
}

func (d Dog) MakeSound() string {
	return "Wow wow"
}

type Cat struct {
	Name     string
	Age      int
	Interest string
}

func (c Cat) String() string {
	return fmt.Sprintf("Cat name: %s, age: %d, interest: %s", c.Name, c.Age, c.Interest)
}

func (c Cat) MakeSound() string {
	return "Meow meow"
}

func generateSound(a Animal) {
	fmt.Println(a.MakeSound())
}

func main() {
	dog := Dog{
		Name:  "Boyka",
		Age:   3,
		Color: "brown",
	}

	cat := Cat{
		Name:     "Masha",
		Age:      2,
		Interest: "Playing with a ball",
	}

	generateSound(dog)
	generateSound(cat)
	fmt.Println(dog)
	fmt.Println(cat)
}
