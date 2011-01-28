package euler

func factorial(n int) int {
    c := 1
    for i := 1; i <= n; i++ { c *= i }
    return c
}

func fsum(n int) int {
    r := 0
    for n != 0 {
        t := n%10
        r += factorial(t)
        n /= 10
    }
    return r
}

func E34() int {
	sum := 0
	for i := 3; i < 1E6; i++ {
		if fsum(i) == i { sum += i }
	}
	return sum
}
