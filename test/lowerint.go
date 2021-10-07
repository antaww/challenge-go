package main

import "fmt"

func main() {
	primes := []int{78, 6, 4}
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
