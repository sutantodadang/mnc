package main

import (
	"bufio"
	"fmt"
	"os"
)

func No3() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("q to quit")

	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "q" {
			break
		}

		if len(text) < 1 || len(text) > 4096 {
			fmt.Println("false")
		}

		allowedChar := map[rune]bool{
			'<': true, '>': true,
			'{': true, '}': true,
			'[': true, ']': true,
		}

		brackerPair := map[rune]rune{
			'>': '<',
			'}': '{',
			']': '[',
		}

		var stack []rune
		var lastopen rune
		flag := true

		for _, v := range text {

			if !allowedChar[v] {
				flag = false
				break
			}

			if v == '<' || v == '{' || v == '[' {

				if lastopen != 0 && lastopen != v {
					flag = false
					break
				}

				stack = append(stack, v)

				lastopen = v
			} else {

				if len(stack) == 0 {

					flag = false
					break
				}
				last := stack[len(stack)-1]

				if brackerPair[v] != last {

					flag = false
					break
				}

				stack = stack[:len(stack)-1]

				if len(stack) == 0 {
					lastopen = 0
				} else {
					lastopen = stack[len(stack)-1]
				}
			}
		}

		if len(stack) == 0 && flag {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}

}
