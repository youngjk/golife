func subtractProductAndSum(n int) int {
    tmpN := n
    product := 1
    sum := 0
    
    for tmpN > 0 {
        digit := tmpN%10
        product *= digit
        sum += digit
        tmpN /= 10
    }
    
    return product - sum
}