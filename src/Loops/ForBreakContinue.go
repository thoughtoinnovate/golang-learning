package main

import "fmt"

func main() {

headerBreakLabel:
	for timer := 10; timer >= 0; timer-- {
		if timer%2 == 0 {
			continue
		}

		if timer == 3 {
			fmt.Println("Break both loops")
			break headerBreakLabel
		}

		fmt.Println(timer)
	}
}
