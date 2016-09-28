package euler

func E14() uint64 {
	const upto = 1000000
	collatz := make([]uint64, upto+1)

	max, ind := uint64(1), uint64(1)
	for i := uint64(1); i <= upto; i++ {
		t, l := i, uint64(1)
		for t != 1 {
			if t%2 == 0 {
				t /= 2
			} else {
				t = 3*t + 1
			}
			if t <= upto && collatz[t] != 0 {
				l += collatz[t]
				break
			}
			l++
		}
		collatz[i] = l
		if l > max {
			max, ind = l, i
		}
	}

	return ind
}
