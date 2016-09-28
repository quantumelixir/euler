package euler

import (
	"github.com/quantumelixir/euler/utils"
)

func E7() (i int) {

	nth := 10001

	for count := 0; count != nth; i++ {
		if utils.IsPrime(i) {
			count++
		}
	}

	return i - 1
}
