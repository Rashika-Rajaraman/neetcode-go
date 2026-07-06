func isPalindrome(s string) bool {
    l, r := 0, len(s) -1

    // Time Complexity - O(n)
    // Space Complexity - O(1)
    for l<r {
        for l<r && !isAlphaNum(rune(s[l])) {
            l++
        }
        for r>l && !isAlphaNum(rune(s[r])) {
            r--
        }
        if unicode.ToLower(rune(s[l])) != unicode.ToLower(rune(s[r])) {
            return false
        }
        l++ 
        r--
    }
    return true
}

func isAlphaNum(ch rune) bool {
    return unicode.IsDigit(ch) || unicode.IsLetter(ch)
}