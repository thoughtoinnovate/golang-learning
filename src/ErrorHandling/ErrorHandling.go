package main

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	//exexute function check if an error
	if val, err := doSomething(); err == nil {
		fmt.Println("Hurray! No error occured val is :", val)
	} else if nil != err {
		fmt.Println("Error occured :", err)
	}

}

func doSomething() (int, error) {
	switch temp := random(); temp {
	case 2, 4, 6, 8:
		fmt.Println("Got even number:", temp)
		return temp, nil
	case 1, 3, 5, 7, 9:
		fmt.Println("got odd number:", temp)
		return temp, nil
	default:
		return -1, errors.New(runtime.GOARCH)
	}
}

/*
Random int generator based on time
*/
func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
