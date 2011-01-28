package euler

func E1() int {
    sum := 0;

    for i := 0; i < 1000; i+=5  { sum += i; }
    for i := 0; i < 1000; i+=3  { sum += i; }
    for i := 0; i < 1000; i+=15 { sum -= i; }

    return sum;
}
