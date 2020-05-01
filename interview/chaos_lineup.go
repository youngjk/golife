package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
    bribeCount := 0

    for i := len(q)-1; i > 0; i-- {
        // If the current value is greater than 2(max) swaps away, lineup is chaotic
        if (int(q[i]) - (i+1) > 2) {
            fmt.Println("Too chaotic")
            return
        }

        // Set the max movable value index
        prevIndex := 0
        if int(q[i]) - 2 > 0 {
            prevIndex = int(q[i]) - 2
        }


        // Check values between new index and current index and increment bribecount when new index value is bigger than current value
        for j := maxIndex; j < i; j++ {
            if q[j] > q[i] {
                bribeCount++
            }
        }
    }

    fmt.Println(bribeCount)
    return
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(readLine(reader), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
    }
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
