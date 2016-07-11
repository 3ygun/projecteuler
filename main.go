package main

import (
	"fmt"
	"strconv"
	"strings"
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
		question1()
	}

	fmt.Println()
}
