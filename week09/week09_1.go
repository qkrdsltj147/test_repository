package main

import "fmt"

func main() {
	var a int = 10
	var pa *int
	pa = &a

	fmt.Println(&a, a)
	fmt.Println(pa, *pa)
	fmt.Println(&pa)
	fmt.Printf("%T %T \n", a, pa)
}
