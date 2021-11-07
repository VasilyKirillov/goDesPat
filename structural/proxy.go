/*
There is a function that returns a sequence of N prime numbers
also there is a Proxy for this function that adds to it a profiling functionality
*/
package main

import (
	"log"
	"time"
)

func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func PrimeNumberSequenceDecorator(N int) (primes []int) {
	start := time.Now()
	primes = sieveOfEratosthenes(N)
	elapsed := time.Since(start)
	log.Printf("It took %s to generate %v prime numbers", elapsed, N)
	return
}
