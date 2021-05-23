func defangIPaddr(address string) string {
    aRune := []rune(address)
    newA := make([]rune, 0)
    
    for i := 0; i < len(address); i++ {
        if address[i] == '.' {
            newA = append(newA, '[')
            newA = append(newA, '.')
            newA = append(newA, ']')
        } else {
            newA = append(newA, aRune[i])
        }
    }
    
    return string(newA)
}