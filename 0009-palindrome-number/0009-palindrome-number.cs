func isPalindrome(x int) bool {
    
    str := strconv.Itoa(x)
    n := len(str)-1
    
    for index, char := range str {    
        if char != rune(str[n-index]){
            return false
        }
    }
    
    return true
}