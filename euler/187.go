package euler

import (
	"github.com/quantumelixir/euler/utils"
)

func E187() int {
	const upto = 1E8
	primes := utils.Primes(upto/2 + 1)

	count := 0
	for i := 1; i < len(primes); i++ {
		for j := i; j < len(primes) &&
			uint64(primes[i])*uint64(primes[j]) < upto; j++ {
			count++
		}
	}

	return count + len(primes) // account for semiprimes of the form 2*p
}
