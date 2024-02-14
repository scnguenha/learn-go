package main

import "fmt"

func main() {
	name := "Sidy"

	fmt.Println(name)
	fmt.Println(&name)
	a := &name

	fmt.Println(*a)

	*a = name + " Nguenha"
	fmt.Println(a)

	fmt.Println(name)
}
