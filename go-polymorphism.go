package main

import "fmt"

type Animal interface {
    Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
    return "Woof"
}

type Cat struct{}

func (c Cat) Speak() string {
    return "Meow"
}

type Person struct {
    Name string
}

func (p *Person) Speak() string {
    return "Hi, my name is " + p.Name
}

type Employee struct {
    Person
    JobTitle string
}

func main() {
    animals := []Animal{Dog{}, Cat{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
e := Employee{Person: Person{Name: "Alice"}, JobTitle: "Engineer"}
    fmt.Println(e.Speak())  // Output: Hi, my name is Alice
    }

}
