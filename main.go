package main

import (
	"fmt"
	"strconv"
	"strings"

	p1to10 "github.com/3ygun/projecteuler/p1to10"
)

func main() {
	input := ""
	fmt.Printf("Enter:\n\t# of question\n\tq to quit\n\n")

	for !(strings.EqualFold(input, "q")) {
		fmt.Printf("Which Question?: ")
		fmt.Scanln(&input)
		qNum, err := strconv.Atoi(input)

		if err == nil {
			runQuestion(qNum)
		} else {
			fmt.Printf(input + " not an question number.\n")
		}

		fmt.Println()
	}
}

func runQuestion(qNum int) {
	fmt.Println()

	switch qNum {
	case 1:
		p1to10.Question1()
	case 2:
		p1to10.Question2()
	}

	fmt.Println()
}
