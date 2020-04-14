package main

import "fmt"

func main() {

	myFavourites := []string{"Java", "Ruby", "Go", "Python"}
	otherFavourites := []string{"Java", "Go", "Python"}

	fmt.Println("Common Fav's below")

	for _, myFav := range myFavourites {
		for _, otherFav := range otherFavourites {
			if myFav == otherFav {
				fmt.Println(myFav)
			}
		}
	}
	fmt.Println("Range loop without underscore prints indexes only")
	for i := range myFavourites {
		fmt.Println(i)
	}

}
