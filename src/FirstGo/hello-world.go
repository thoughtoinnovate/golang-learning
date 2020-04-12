package main

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
)

//Defining constants
const (
	income = 1000000
)

//package level variables
var (
	name, class string = "my name is good", "one"
	rollNumber  int
	module      = 2.4

	//type infer
	x, y, z = "this is x ", 'y', "this is z"
	d       = 10
)

/*
 */
func main() {
	fmt.Println("Jai Mata Di!")
	fmt.Println("Welcome my os name", runtime.GOOS)
	fmt.Println("Print variables below")
	fmt.Println("name:", name, reflect.TypeOf(name))
	fmt.Println("roll:", rollNumber, reflect.TypeOf(rollNumber))
	fmt.Println("module", module, reflect.TypeOf(module))
	fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y), reflect.TypeOf(z))

	//Adding variable of different types
	a := 1
	b := 2.01
	c := a + int(b)
	//local declaration with :=
	d := 3
	fmt.Println("the c is:", c)
	fmt.Println(d)

	//pointers
	ptr := &d
	fmt.Println("the pointer address of d is ", ptr)
	fmt.Println("The pointer value of d is:", *ptr)

	//:= for declaring variable in scope
	//= for reassignment within the scope

	//Go is pass by value
	fmt.Println("Check go is pass by value original name:", name)
	//change name by calling function
	ChangeName(name)
	fmt.Println("After name change function", name)

	//Changing the actual value of the variable
	ChangeNameActual(&name)
	fmt.Println("Check if actual name is changed:", name)

	//printing
	fmt.Println("The Constatnt income is : ", income)

	//print environment variables
	//environmentVars := os.Environ()
	fmt.Println("Environment Variables below =======================>")
	/* for _, env := range os.Environ() {
		fmt.Println(env)
	}*/
	// print environment logged in user
	UsrName := os.Getenv("USER")
	fmt.Println("The logged in user is :", UsrName)

}

func ChangeName(localName string) string {
	localName = "changed name"
	fmt.Println("name changed inside func is :", localName)
	return localName
}

func ChangeNameActual(localName *string) string {
	*localName = "changed name"
	fmt.Println("name changed inside actual func is :", *localName)
	return *localName
}
