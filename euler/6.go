package euler

func E6() int {
	suma, sumb := 0, 0
	for i := 1; i <= 100; i++ {
		suma += i
		sumb += i * i
	}
	return suma*suma - sumb
}
