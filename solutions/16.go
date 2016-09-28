package euler

import (
	"github.com/quantumelixir/euler/utils"
)

type BigInt []int // base 10

func lmultiply(a BigInt, b BigInt) BigInt {
	carry, orig := 0, len(a)
	for i := 0; ; i++ {
		if i == len(a) {
			a = a[:len(a)+1]
		}
		a[i] = a[i] * b[0]
		for j := 0; j < i; j++ {
			if j < len(a) && i-j < len(b) {
				a[i] += a[j] * b[i-j]
			}
		}
		if i >= orig-1 && a[i]+carry == 0 {
			break
		}
		carry += a[i]
		a[i] = carry % 10
		carry /= 10
	}

	return a
}

func E16() int {
	res := make([]int, 1, 1000)
	base := make([]int, 1)
	res[0] = 2
	base[0] = 2

	for i := 2; i <= 1000; i++ {
		res = lmultiply(res, base)
	}
	return utils.Sum(res)
}
