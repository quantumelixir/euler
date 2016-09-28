package euler

import (
//	"fmt"
	"math/big"
)

func E469() string {
	s, pq, q := big.NewRat(0, 1), big.NewRat(1, 1), big.NewRat(1, 1)
	oans, cans := big.NewRat(0, 1), big.NewRat(1, 1)
	n, prec, t := int64(3), big.NewRat(1, 1e8), big.NewRat(0, 1)
	prec.Mul(prec, prec) // precision is 1e-16
	for t.Abs(t.Sub(oans, cans)).Cmp(prec) == 1 {

		s.Add(s, pq)

		pq.Set(q)

		t.Quo(big.NewRat(2, 1), big.NewRat(n, 1))
		t.Mul(t, s)
		t.Add(t, big.NewRat(1, 1))
		q.Set(t)

		t.Add(q, big.NewRat(1, 1))
		t.Quo(t, big.NewRat(n+3, 1))
		t.Sub(big.NewRat(1, 1), t)
		oans.Set(cans)
		cans.Set(t)

		// fmt.Println("n:", n, "ans(n+3):", cans.FloatString(14))
		n++
	}
	return cans.FloatString(14)
}
