package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
		var valleyCount int32 = 0
    currentLevel := 0
		isBelowSeaLevel := false

    for i := 0; i < int(n); i++ {
			// Before processing the input make sure we determine if we are below/above sea level
			if(currentLevel < 0 && !isBelowSeaLevel) {
					isBelowSeaLevel = true
			}
			// Compute current level
			if(s[i] == 'D') {
					currentLevel--
			} else if (s[i] == 'U') {
					currentLevel++
			}
			// Check if we have a valley
			if(isBelowSeaLevel && currentLevel >= 0) {
					valleyCount++
					isBelowSeaLevel = false
			}
    }

    return valleyCount
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    s := readLine(reader)

    result := countingValleys(n, s)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
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
