package main

import "fmt"

func main() {
	productMap := createMap()
	iterateMap(productMap)
	modifyMap(productMap)
}

func createMap() map[string]int {
	productCountMap := make(map[string]int)
	productCountMap["Mobiles"] = 10
	productCountMap["PC"] = 100
	productCountMap["TV Units"] = 20

	fruitCountMap := map[string]int{"Apple": 7, "Mango": 5}

	fmt.Printf("Product map created using make:%v\nFruit Map created dynamically:%v\n", productCountMap, fruitCountMap)
	return productCountMap
}

func iterateMap(mapType map[string]int) {

	fmt.Println("Iterating maps and printing below:")

	for key, val := range mapType {
		fmt.Printf("key:%s Value:%d\n", key, val)
	}
}

func modifyMap(mapType map[string]int) {

	fmt.Println("Modifying Map Below..")
	fmt.Println("Original Map:", mapType)

	//modifying values in map
	mapType["PC"] = 500
	fmt.Println("Modified value of map in PC :", mapType)

	//inserting values in map
	mapType["New Product"] = 101
	fmt.Println("Map after inserting new value:", mapType)

	//Deleting value from the map
	delete(mapType, "New Product")
	fmt.Println("Map after deleting new value:", mapType)

}
