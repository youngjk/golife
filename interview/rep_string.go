package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
    // If there is no a return 0
    if(strings.Count(s, "a") == 0) {
        return 0;
    }
    // Get count of S in n and remainder
    countOfS := int(n)/len(s)
    remString := int(n)%len(s)

    // Calculate count of a
    countOfA := strings.Count(s, "a") * countOfS
    if(strings.Count(s[0:remString], "a") > 0) {
        countOfA += strings.Count(s[0:remString], "a")
    }

    return int64(countOfA)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    s := readLine(reader)

    n, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)

    result := repeatedString(s, n)

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
