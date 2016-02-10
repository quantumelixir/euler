package euler

import (
	"github.com/quantumelixir/euler/utils"
)

func E85() int {
	const limit = 2000000
	min, mini, minj := limit, 0, 0

	// upper limit of 4000 obtained by solving the quadratic equation
	// i*(i+1)/2*j(j+1)/2 = 2000000 for i with j = 1; (ix1 grid)
	for i := 1; i <= 4000; i++ {
		for j := 1; j <= i; j++ {
			curr := i * (i + 1) / 2 * j * (j + 1) / 2
			if utils.Abs(curr-limit) < min {
				min, mini, minj = utils.Abs(curr-limit), i, j
			}
			if curr > limit {
				continue
			}
		}
	}

	return mini * minj
}
