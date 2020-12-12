package main

import "fmt"

func main() {
	var primes [] int
	var i int = 0

	for i=0; i<7; i++ {
		primes = append(primes, i)
	}

fmt.Println(primes)
}