package main

import "fmt"

func isIncreasing(n uint64) bool {
	var max uint64
    for {
		max = n % 10
		n /= 10
		if n == 0 {
			break
		}
		if n%10 > max {
			return false
		}
    }
	return true
}

func isDecreasing(n uint64) bool {
	var min uint64
    for {
		min = n % 10
		n /= 10
		if n == 0 {
			break
		}
		if n%10 < min {
			return false
		}
    }
	return true
}

func isBouncy(n uint64) bool {
    return !(isIncreasing(n) || isDecreasing(n))
}

func main() {
	var bnc, i uint64;
	for i = 1; ; i++ {
		if isBouncy(i) { bnc++ }
		if bnc*100 == 99*i {
			fmt.Println(i)
			break
		}
	}
}
