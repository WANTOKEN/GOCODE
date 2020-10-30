package main

import "fmt"

type Person struct {
    age int
}

func main() {
    person := &Person{28}

    // 1. 
    defer fmt.Printf("%s:%d","1",person.age)

    // 2.
    defer func(p *Person) {
        fmt.Printf("%s:%d","2",p.age)
    }(person)  

    // 3.
    defer func() {
        fmt.Printf("%s:%d","3",person.age)
    }()

	person.age = 29
}