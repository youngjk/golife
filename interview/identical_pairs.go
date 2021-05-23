func numIdenticalPairs(nums []int) int {
    pairCount := 0
    lookupTable := make([]int, 101)
    
    for _, v := range nums {
        pairCount += lookupTable[v]
        lookupTable[v]++
    }
    
    return pairCount
}