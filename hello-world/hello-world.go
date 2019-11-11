// Golang packages modules
package main

// fmt - go print
import "fmt"

func main() {
	const name, language = "YJ", "Golang"
	var printType string

	fmt.Println("Print Type: Print / Printf / Println")
	if _, err := fmt.Scan(&printType); err == nil {
		switch printType {
		case "Print":
			// Variables are directly placed inside
			fmt.Print(printType + ": Hello " + name + " Welcome to " + language + "\n")
		case "Printf":
			// Variables are defined after string
			fmt.Printf("%s: Hello %s Welcome to %s \n", printType, name, language)
		case "Println":
			// Spaces are appended between each params + newline appended
			fmt.Println(printType+":", "Hello", name, "Welcome to", language)
		default:
			fmt.Println("INVALID INPUT PLEASE ENTER: Print / Printf / Println")
		}
	}
}

// go run <file.go> => Runs the file
// go build <file.go> => Builds executable
