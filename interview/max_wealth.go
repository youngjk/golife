func maximumWealth(accounts [][]int) int {
    // Iterate through and figure out the max sum of each row
    // If new determined value is bigger make that max
    max := 0
    curWealth := 0
    
    for _, a := range accounts {
        for _, b := range a {
            curWealth += b
        }
        if curWealth > max {
            max = curWealth
        }
        curWealth = 0
    }
    
    return max
}