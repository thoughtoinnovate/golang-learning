package main

import "fmt"

func main() {
	//make slice slice type,length,capacity
	mySlice := make([]int, 5, 10)

	fmt.Printf("The length of slice is: %d and capacity of slice is: %d\n", len(mySlice), cap(mySlice))

	//shortcut
	fmt.Println("Shortcut for creating slice")
	shortSlice := []string{"a", "b", "c"}
	fmt.Printf("The length of slice is: %d and capacity of slice is: %d\n", len(shortSlice), cap(shortSlice))

	//Accessing slices
	accessAndChangValueOfSlices()

	//Appending and verifying growable slice backed by an array
	appendToSlice()
}

func accessAndChangValueOfSlices() {

	intSlice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("original int array at index 1:", intSlice[1])
	//modify
	intSlice[1] = 1
	fmt.Println("modified int array at index1:", intSlice[1])

	//custom slice include from index 2 till index 4 , 5 is excluded
	sliceOfSlice := intSlice[2:5]
	fmt.Println(sliceOfSlice)

	//Create lsice without start index , it starts from 0 index
	sliceFromstart := intSlice[:3]
	fmt.Println("sliceFromstart:", sliceFromstart)
	//Create Slice without end index , it goes till last inde
	sliceFromEnd := intSlice[3:]
	fmt.Println("sliceFromEnd:", sliceFromEnd)
}

func appendToSlice() {
	myslice := make([]int, 2, 2)
	fmt.Printf("This length of slice is:%d\nThe capacity of slice is:%d\n", len(myslice), cap(myslice))

	//append it will double the capacity and lenghth of underlying array once the last index is added to the capacity
	for i := 0; i < 2; i++ {
		myslice = append(myslice, i+4)
	}

	fmt.Printf("This length of slice is:%d\nThe capacity of slice is:%d\n", len(myslice), cap(myslice))

	fmt.Println("append slice to slice")
	fmt.Println("before appended:", myslice)
	newSlice := []int{10, 11}
	myslice = append(myslice, newSlice...)
	fmt.Println("Slices appended:", myslice)

}
