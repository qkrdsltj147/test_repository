package main

import "fmt"

func main() {
	//var primes [3]int
	//primes[0] = 2
	//primes[2] = 5

	primes := [3]int{2, 3, 5} //initialize by array literal

	for i := 0; i < 3; i++ {
		fmt.Println(primes[i])
	}
}
