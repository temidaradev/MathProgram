package rectangle

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Rectangle() {
	var pl = fmt.Println
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		pl("Welcome...")

		pl("Area or Perimeter")
		scanner.Scan()
		secim := strings.ToLower(scanner.Text())

		if secim == strings.ToLower("q") {
			break
		}

		pl("Type Long Side Of Rectangle")
		scanner.Scan()
		long, _ := strconv.ParseFloat(scanner.Text(), 64)

		pl("Type Short Side Of Rectangle")
		scanner.Scan()
		short, _ := strconv.ParseFloat(scanner.Text(), 64)

		if short > long {
			pl("Please Obey My Rules!!!")
			break
		} else {
			pl("Procces Completed Sucessfully!!!")
		}

		if secim == "area" {
			fmt.Println(long * short)
		} else if secim == "perimeter" {
			fmt.Println((long + short) * 2)
		}
	}
}
