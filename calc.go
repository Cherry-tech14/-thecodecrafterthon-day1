package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func validAction(s string) bool {
	return s == "add" || s == "sub" || s == "mul" || s == "div" || s == "HELP" || s == "quit"
}
func main() {

	blue := "\033[34m"
	red := "\033[31m"
	green := "\033[32m"
	cherry := "\033[35m"
	yellow := "\033[33m"
	reset := "\033[0m"

	for {
		fmt.Println(green + "Welcome to my simple CLI calculator!" + reset)

		fmt.Println(cherry + "Enter action" + reset)
		function := bufio.NewScanner(os.Stdin)
		function.Scan()
		input := function.Text()

		part := strings.Fields(input)
		action := part[0]
		if validAction(action) {
			switch action {
			case "add":
				if len(part) != 3 {
					fmt.Println(red + "Improper action" + reset)
					continue
				}
				num1, err := strconv.Atoi(part[1])
				if err != nil {
					fmt.Println(red + "wrong input" + reset)
					continue
				}
				num2, err := strconv.Atoi(part[2])
				if err != nil {
					fmt.Println(red + "wrong input" + reset)
					continue
				}
				output := num1 + num2
				fmt.Println(blue+"result :"+reset, output)
			case "sub":
				if len(part) != 3 {
					fmt.Println(red + "Improper action" + reset)
					continue
				}
				num1, err := strconv.Atoi(part[1])
				if err != nil {
					fmt.Println(red + "wrong input" + reset)
					continue
				}
				num2, err := strconv.Atoi(part[2])
				if err != nil {
					fmt.Println(red + "wrong input" + reset)
					continue
				}
				output := num1 - num2
				fmt.Println(blue+"result :"+reset, output)
			case "mul":
				if len(part) != 3 {
					fmt.Println(red + "Improper action" + reset)
					continue
				}
				num1, err := strconv.Atoi(part[1])
				if err != nil {
					fmt.Println(red + "wrong input" + reset)
					continue
				}
				num2, err := strconv.Atoi(part[2])
				if err != nil {
					fmt.Println(red + "wrong input" + reset)
					continue
				}
				output := num1 * num2
				fmt.Println(blue+"result :"+reset, output)
			case "div":

				num1, err := strconv.Atoi(part[1])
				if err != nil {
					fmt.Println(red + "wrong input" + reset)
					continue
				}
				num2, err := strconv.Atoi(part[2])
				if err != nil {
					fmt.Println(red + "wrong input" + reset)
					continue
				}

				if num2 == 0 {
					fmt.Println(red + "cannot divide by 0" + reset)
					continue
				} else {
					output := num1 / num2
					fmt.Println(blue+"result :"+reset, output)
				}

			case "HELP":
				fmt.Println(yellow + "The program supports the following commands: \naddition       (Usage: add <a> <b>)\nmultiplication (Usage: mul <a> <b>)\nsubtraction    (Usage: sub <a> <b>)\ndivision       (Usage: div <a> <b>)" + reset)
				continue

			}
			if action == "quit" {
				fmt.Println(green + "Goodbye" + green)
				break
			}

		} else if !validAction(action) {
			fmt.Println(cherry + "Do you need help?\nIf yes, Enter help" + reset)
			continue
		}

	}

}
