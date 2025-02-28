package main

import (
	"fmt"
)

// the main function to start the process
func _() {
	input := "This is a test input. All the data is not relevant. The facts described are fictional."
	fmt.Println(input)
}

func main() {
	categories := GetDataset("./categories.json")
	fmt.Println(categories)
}
