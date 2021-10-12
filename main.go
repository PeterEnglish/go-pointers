package main

import "fmt"

func main() {
	fmt.Println("This is all about pointers")

	number1 := 2
	number3 := &number1
	*number3 = 45
	fmt.Println(number3)
	fmt.Println(number1)
	//this will print out the memory address of the value - a pointer
	//go is a garbage collected language
	//Pointers can have a benefit when you have structs containing lots of
	//data. When you have these, the overhead of the garbage collector might
	//be negated by the overhead you’d get when copying large amounts of data.
	//Reasons to use pointer:
	//The only way to mutate a variable that you pass to a function is
	//by passing a pointer
	//WHen you do not do this, you pass a copy and not a reference to the origina value.

	//EG
	createPerson()            //Still print out Richard, as we do not change original value
	createPersonWithPointer() //This works
}

type person struct {
	name string
}

func createPerson() {
	p := person{"Richard"}
	rename(p)
	fmt.Println(p)
}

func rename(p person) {
	p.name = "test"
}

func createPersonWithPointer() {
	p := person{"Richard"}
	renameWithPointer(&p)
	fmt.Println(p)
}

func renameWithPointer(p *person) {
	p.name = "test"
}
