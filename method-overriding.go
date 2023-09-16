package main

import "fmt"

// Animal interface
type Animal interface {
	Speak() string
}

// Dog struct
type Dog struct{}

// Speak method for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Cat struct
type Cat struct{}

// Speak method for Cat, essentially "overriding" the Speak method in Animal
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	animals := []Animal{
		Dog{},
		Cat{},
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
