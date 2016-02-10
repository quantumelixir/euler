package euler

import (
	"github.com/quantumelixir/euler/utils"
)

func zpow(n int) (z int) {
	for z = 1; z <= n; z *= 10 {
	}
	return z
}

func E38() (max int) {

	for i := 1; i < 1E5; i++ {
		var used [10]bool
		cat, count, flag := 0, 0, true
	J:
		for j := 1; j <= 9; j++ {
			n := i * j
			for count < 9 && n != 0 {
				if used[n%10] {
					flag = false
					break J
				} else {
					used[n%10] = true
					count++
				}
				n /= 10
			}
			if n == 0 && count == 9 && used[0] == false {
				for k := 1; k <= j; k++ {
					cat = cat*zpow(i*k) + i*k
				}
				break
			}
		}
		if flag {
			max = utils.Max(cat, max)
		}
	}

	return max
}
