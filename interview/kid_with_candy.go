func kidsWithCandies(candies []int, extraCandies int) []bool {
    // Determine max value in array
    // Iterate through to see which candies[i] > max-extraCandies
    isMaxArray := make([]bool, len(candies))
    max := 0
    for _, v := range candies {
        if v > max {
            max = v
        }
    }
    
    threshold := max - extraCandies
    for i, cc := range candies {
        if cc >= threshold {
            isMaxArray[i] = true
        }
    }
    
    return isMaxArray
}