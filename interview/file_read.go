package main

import (
	"fmt"
	"io/ioutil"
	"strings"
  // "strconv"
)

type Logs struct {
  user string
  pid string
  cpu string
  mem string
  time string
  command string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var err error
	logArray := make([]Logs, 0)

	// Read file
	fileRead, err := ioutil.ReadFile("./practice.txt")
	check(err)
	lineRead := strings.Split(string(fileRead), "\n")

	for i, v := range lineRead {
    // Skip headline
    if i == 0 {
      continue
    }
    // Create Object
		fieldRead := strings.Fields(v)
    logObject := Logs{
      user: fieldRead[0],
      pid: fieldRead[1],
      cpu: fieldRead[2],
      mem: fieldRead[3],
      time: fieldRead[9],
      command: fieldRead[10],
    }
		logArray = append(logArray, logObject)
	}
  fmt.Println(logArray)
}