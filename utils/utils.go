package utils

func Abs(a int) (int) {
    if a > 0 { return a }
    return -a
}

func AbsSum(a []int) (int) {
    sum := 0
    for _, i := range a { sum += Abs(i) }
    return sum
}

func Sum(a []int) (int) {
    sum := 0
    for _, i := range a { sum += i }
    return sum
}

func Min(a, b int) (int) {
    if a < b {
        return a
    }

    return b
}

func Max(a, b int) (int) {
    if a > b {
        return a
    }

    return b
}

func Max64(a, b int64) (int64) {
    if a > b {
        return a
    }

    return b
}

func RotateInt(n int) (int) {

    sign := 1
    if n < 0 {
        sign = -1
        n = -n
    }

    r := n % 10
    p := 1
    for p < n {
        p *= 10
    }

    return sign*(r*p + n)/10
}

func ReverseInt(n int) (int) {
    var res int
    for n != 0 {
        res = res*10 + n%10
        n /= 10
    }
    return res
}

func IsPrime(n int) (bool) {

    if n == 2 || n == 3 {
        return true
    }

    if n == 1 || n % 2 == 0 || n % 3 == 0 {
        return false
    }

    sqrt := 1
    for sqrt*sqrt < n {
        sqrt++
    }

    for j := 5; j <= sqrt; j += 6 {
        if n%j == 0 || n%(j+2) == 0 {
            return false
        }
    }

    return true
}

//
// Return a boolean slice with false in positions whose indices are primes,
// upto (but not including) n.
// 
func Sieve(n int) (sieve []bool) {
    sieve = make([]bool, n)
    sieve[0], sieve[1] = true, true
    for i := 2; i < n; i++ {
        for j := i*2; !sieve[i] && j < n; j += i {
            sieve[j] = true
        }
    }
    return sieve
}

func Primes(n int) (primes []int) {

    count, sieve := 0, Sieve(n)
    for _, i := range sieve { if !i { count++ } }

    k, primes := 0, make([]int, count)
    for i, j := range sieve {
        if !j {
            primes[k] = i
            k++
        }
    }
    return primes
}

func ArePermutations(a, b int) (bool) {

    list := make([]int, 10)

    for a != 0 {
        list[a % 10]++
        a /= 10
    }
    for b != 0 {
        list[b % 10]--
        b /= 10
    }

    return AbsSum(list) == 0
}

func Pow(n, k int) int {
    c := 1
    for i := 1; i <= k; i++ { c *= n }
    return c
}

func GCD(a, b int) int {
    if a < 0  { return GCD(-a, b) }
    if b < 0  { return GCD(a, -b) }
    if a < b  { return GCD(b, a) }
    if b == 0 { return a }
    return GCD(b, a%b)
}

func NCR(n, r int) int {
    if n < r { return 0 }
    c := 1
    for i := 1; i <= r; i++ {
        c = c*(n - i + 1)/i
    }
    return c
}

func DigitSum64(n uint64) uint64 {
    sum := uint64(0)
    for n != 0 {
        sum += n%10
        n /= 10
    }
    return sum
}

func DigitSum(n int) int {
    sum := 0
    for n != 0 {
        sum += n%10
        n /= 10
    }
    return sum
}
