package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Below excution of if statements")
	learnIf()
	fmt.Println("Learn Simple Switch below")
	learnSimpleSwitch("one", "two", "three", "something else")
	fmt.Println("Learn below Switch with fall through")
	learnFallThroughSwitch("TWO")
	fmt.Println("Learn below for multiple case handling in single swicth case")
	learnMultiCaseSwitch()
}

func learnIf() {
	//if with initialization
	if first, second := 30, 40; first > second {
		fmt.Println("First is greate")
		if first > 100 {
			fmt.Println("first made century")
		}
	} else if second > first {
		fmt.Println("Second is greate")

	} else {
		fmt.Println("Both are equal")
	}
}

func learnSimpleSwitch(args ...string) {
	for _, val := range args {
		switch strings.ToLower(val) {
		case "one":
			fmt.Println("One Executed")
		case "two":
			fmt.Println("Second Executed")
		case "three":
			fmt.Println("three Executed")
		default:
			fmt.Println("Some other val passed")
		}
	}
}

/*
This  function tells how to ver ride implicit break added by go switch with fallthrough keyword
Default isnt executed in fall through
*/
func learnFallThroughSwitch(val string) {

	switch strings.ToLower(val) {
	case "one":
		fmt.Println("One Executed")
	case "two":
		fmt.Println("Second Executed")
		fallthrough
	case "three":
		fmt.Println("three Executed")
	default:
		fmt.Println("Some other val passed")
	}

}

func learnMultiCaseSwitch() {
	switch temp := random(); temp {
	case 2, 4, 6, 8:
		fmt.Println("Got even number:", temp)
	case 1, 3, 5, 7, 9:
		fmt.Println("got odd number:", temp)
	default:
		fmt.Println("got some other number:", temp)
	}
}

/*
Random int generator based on time
*/
func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
