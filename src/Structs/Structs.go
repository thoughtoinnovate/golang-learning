package main

import "fmt"

func main() {
	//create a struct
	type Book struct {
		Author string
		Genre  string
		price  int
	}

	//Create Object of struct way 1
	var CodingBook = new(Book)
	//If we print it we should get default values , by default values are zero
	fmt.Println(CodingBook)

	//Way 2 initialize variables
	StoryBook := Book{
		Author: "Thought",
		Genre:  "Story",
		price:  900,
	}

	fmt.Println(StoryBook)
	fmt.Println("Author of Book is:", StoryBook.Author)
	//change price of the book
	StoryBook.price = 500
	fmt.Println("Storu Book after discount is:", StoryBook.price)

}
