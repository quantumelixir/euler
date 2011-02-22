package euler

func isDecimalPalindrome(n uint64) bool {
    m, t := uint64(0), n
    for t != 0 {
        m *= 10
        m += t % 10
        t /= 10
    }
    return m == n
}

func isBinaryPalindrome(n uint64) bool {
    m, t := uint64(0), n
    for t != 0 {
        m <<= 1
        m += t & 0x1
        t >>= 1
    }
    return m == n
}

func E36() (sum uint64) {

    var tp, i uint64

    // binary palindromes must be odd
    for tp, i = 1, 1; i < 1E6; i+=2 {
        if i > 10*tp { tp *= 10 }
        if i/tp != i%10 { continue }
        if isBinaryPalindrome(i) && isDecimalPalindrome(i) {
            sum += i
        }
    }

    return sum
}
