package euler

import (
	"github.com/quantumelixir/euler/utils"
	"math"
)

func E3() uint64 {

	n, res := uint64(600851475143), uint64(0)
	nf := float64(n)
	primes := utils.Primes(int(math.Sqrt(nf)))
	for _, p := range primes {
		for n%uint64(p) == 0 {
			n /= uint64(p)
			res = uint64(p)
		}
	}
	return res
}
