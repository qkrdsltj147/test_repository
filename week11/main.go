package main

import "fmt"

func main() {
	primes := [3]int{2, 3, 5} //initialize by array literal

	for i := 0; i < 3; i++ {
		fmt.Println(primes[i])
	}

	//초기화 하지 않은 원소의 제로 값은 해당 원소 타입의 제로값으로 결정된다.

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) //zerovalue
	fmt.Println(test)

	fmt.Printf("%#v\n", primes)
	fmt.Printf("%#v\n", test)

	//fmt.Println(test[5]) 컴파일 에러 our of range

	i := 0
	for i < len(test) { //while
		fmt.Println(test[i])
		i++
	}

	//for idx, prime := range primes { //index만 출력
	//	fmt.Println(idx, prime)
	//}
	for _, prime := range primes {
		fmt.Println(prime)
	}
}
