package main

import (
	"Go/Calculator"
	circle "Go/Cricle"
	"Go/Input"
	rectangle "Go/Rectangle"
	"bufio"
	"fmt"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for true {
		Input.Input()

		fmt.Println("Geometry or Calculator?")
		scanner.Scan()
		gc := scanner.Text()

		if gc == strings.ToLower("geometry") {
			geometry()
		} else if gc == strings.ToLower("calculator") {
			calculator()
		}
	}
}
func geometry() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Circle or Rectangle?")
	scanner.Scan()
	cr := scanner.Text()
	if cr == strings.ToLower("circle") {
		circle.Cricle()
	} else if cr == strings.ToLower("rectangle") {
		rectangle.Rectangle()
	}
	fmt.Println("If Your Transactions End Just Press \"q\"")
}
func calculator() {
	Calculator.Calculator()
}
