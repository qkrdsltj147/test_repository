package main

import "fmt"

func double(number *int) {
	*number *= 2
}

func main() {
	var amount int = 5
	double(&amount)
	fmt.Println("%d \n", amount)
}
