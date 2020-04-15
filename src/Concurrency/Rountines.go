package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	//Executes go routines parallel on multi core cpu , argument 2 means 2 core will
	//be used for parallel processing.
	runtime.GOMAXPROCS(2)
	//create wait groups , means it make main routine to wait for sub routine
	var waitGrp sync.WaitGroup

	fmt.Println("Start execution..")
	//Count in Add arguments , it becomes 0 all routines are released , here we have two routines
	//so waiting for two routines
	waitGrp.Add(2)

	//Anonymous functions
	//'go' keword make them go rountine
	go func() {

		//waits until routines work is done
		defer waitGrp.Done()

		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	//Anonymous functions
	//'go' keword make them go rountine
	go func() {
		//waits until routines work is done
		defer waitGrp.Done()

		fmt.Println("Bye")
	}()

	waitGrp.Wait()
	fmt.Println("Execution Complete")
}
