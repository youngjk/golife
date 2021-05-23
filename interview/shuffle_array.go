func shuffle(nums []int, n int) []int {
    resultArray := make([]int, 2*n)
    
    for i:=0; i<n; i++ {
        resultArray[i*2] = nums[i]
        resultArray[i*2+1] = nums[n+i]
    }
    return resultArray
}