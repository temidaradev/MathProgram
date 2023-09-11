package circle

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)


func Cricle() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Println("Welcome...")

		fmt.Println("Area Or Circumference?")
		scanner.Scan()
		secim := strings.ToLower(scanner.Text())

		if secim == strings.ToLower("q") {
			break
		}

		fmt.Println("Please Enter Radius Of The Circle")
		scanner.Scan()
		radius, _ := strconv.ParseFloat(scanner.Text(), 64)

		if secim == "area" {
			fmt.Println(math.Pi * radius * radius)
		} else if secim == "circumference" {
			fmt.Println(math.Pi * 2 * radius)
		} else {
			fmt.Println("Wrong Entry")
		}
	}

}
