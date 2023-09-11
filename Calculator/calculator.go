package Calculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Calculator() {
	scanner:=bufio.NewScanner(os.Stdin)
	for{
		fmt.Println("1-Addition\n2-Subtraction\n3-Multiplication\n4-Division")
		scanner.Scan()
		secim := scanner.Text()
		if secim == strings.ToLower("q"){
			break
		}
		fmt.Print("Enter First Number:")
		scanner.Scan()
		sayi1, _ := strconv.ParseFloat(scanner.Text(), 64)
		fmt.Print("Enter Second Number:")
		scanner.Scan()
		sayi2, _ := strconv.ParseFloat(scanner.Text(), 64)
		if secim == "1" {
			fmt.Printf("%.2f + %.2f = %.2f", sayi1, sayi2, sayi1+sayi2)
		} else if secim == "2" {
			fmt.Printf("%.2f - %.2f = %.2f", sayi1, sayi2, sayi1-sayi2)
		} else if secim == "3" {
			fmt.Printf("%.2f * %.2f = %.2f", sayi1, sayi2, sayi1*sayi2)
		} else if secim == "4" {
			fmt.Printf("%.2f / %.2f = %.2f", sayi1, sayi2, sayi1/sayi2)
		} else {
			fmt.Println("Wrong Entry")
		}
		fmt.Println("If Your Transactions End Just Press \"q\"")
	}
}
