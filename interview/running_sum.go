func runningSum(nums []int) []int {
    rSum := make([]int, len(nums))
    rSum[0] = nums[0]
    
    for i := 1; i < len(nums); i++ {
        rSum[i] = nums[i] + rSum[i-1]
    }
    
    return rSum
}