package euler

import (
	"github.com/quantumelixir/euler/utils"
)

func isPalindrome(n int) bool {
	return n == utils.ReverseInt(n)
}

func E4() int {

	var res int
	for i := 100; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			if isPalindrome(i * j) {
				res = utils.Max(res, i*j)
			}
		}
	}

	return res
}
