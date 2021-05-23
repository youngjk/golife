func numJewelsInStones(jewels string, stones string) int {
    // How many chars in stones are in jewels
    // Hashmap of jewels
    // Iterate over stones stringe char by char and increment count when hash
    jCount := 0
    
    for _, v:= range stones {
        if strings.Contains(jewels, string(v)) {
            jCount++
        }
    }
    
    return jCount
}