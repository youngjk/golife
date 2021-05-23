func smallerNumbersThanCurrent(nums []int) []int {
    // Sort Array
    // Create Hashmap that maps number to index
    // Iterate through numbs and draw the smaller number by index position
    lookup := [101]int{0}
    result := make([]int, len(nums))
    
    for _, v := range nums {
        lookup[v]++
    }
    
    for j := 0; j < len(nums); j++ {
        for k := 0; k < nums[j]; k++ {
            result[j] += lookup[k]
        }
    }
    
    return result
}
