// Ported from Daniel Scocco's http://www.programminglogic.com/the-sieve-of-eratosthenes-implemented-in-c/

package main

import "fmt"

func main() {
	const LIMIT = 1500000
	const PRIMES = 100000
	var i, j int
	var numbers [LIMIT]int
	var primes [PRIMES]int

	for i = 0; i < LIMIT; i++ {
		numbers[i] = i + 2
	}

	/* sieve the non-primes */
	for i = 0; i < LIMIT; i++ {
		if numbers[i] != -1 {
			for j = 2*numbers[i] - 2; j < LIMIT; j += numbers[i] {
				numbers[j] = -1
			}
		}
	}

	/* transfer the primes to their own array */
	j = 0
	for i = 0; i < LIMIT && j < PRIMES; i++ {
		if numbers[i] != -1 {
			primes[j] = numbers[i]
			j++
		}
	}

	for i = 0; i < PRIMES; i++ {
		fmt.Printf("%d\n", primes[i])
	}

}
