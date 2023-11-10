package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{Name: "Alice", Age: 21},
		{Name: "Bob", Age: 22},
	}

	for _, person := range people {
		fmt.Println(person.Name, "is", person.Age, "years old")
	}
}
