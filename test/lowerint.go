package main

import "fmt"

func main() {
	primes := []int{4, 6, 10}
	for i := 0; i < len(primes)-1; i++ {
		if primes[i] < primes[len(primes)-1] {
			fmt.Println(primes[i])
			break
		}
		if primes[i] > primes[len(primes)-1] {
			fmt.Println(primes[len(primes)-1])
			break
		}
	}
}
