package euler

import (
	"fmt"
)

func b(n uint) float64 {
	m := make(map[uint]float64)
	m[0] = 0
	var rec func(uint) float64
	rec = func(n uint) float64 {
		v, ok := m[n]
		if ok { return v }
		g := (1+n)/2
		m[n] = float64(1.0) + rec(g-1)*float64(g-1)/float64(n) + rec(n-g)*float64(n-g)/float64(n)
		return m[n]
	}
	return rec(n)
}

// Q[n_] := (1/n) (2 n - 1 + (n + 1) Q[n - 1]);
// Q[1] := 1;
// Q[2] := 3;
// R[n_] := Q[n]/n;

func r(n uint) float64 {
	if n == 0 { return 0 }
	if n == 1 { return 1 }
	pq, q := float64(1.0), float64(3.0)
	for i := uint(3); i <= n; i++ { // TOO SLOW!
		pq = q
		q = (-1.0 + float64(2*i) + float64(i+1)*pq)/float64(i)
	}
	return q/float64(n)
}

func E527() string {
	return fmt.Sprintf("%.10g", r(1E10) - b(1E10))
}
