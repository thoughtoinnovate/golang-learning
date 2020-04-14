package main

import (
	"fmt"
)

func main() {

	//simple for loop implementation
	learnSimpleForLoop()
}

func learnSimpleForLoop() {
	for timer := 10; timer >= 0; timer-- {
		if timer == 0 {
			fmt.Println("Boom its time has reached")
			break
		}
		fmt.Println(timer)
		//time.Sleep(1*time.Second)
	}
}
