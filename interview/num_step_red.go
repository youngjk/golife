func numberOfSteps(num int) int {
    tmpN := num
    numSteps := 0
    
    for tmpN > 0 {
        if tmpN%2 == 0 {
            tmpN /= 2
        } else {
            tmpN -= 1
        }
        numSteps++
    }
    
    return numSteps
}