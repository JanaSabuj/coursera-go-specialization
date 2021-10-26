package main

import "fmt"

func main() {
	var fl_num float64

	fmt.Printf("Enter a floating number: ")
	fmt.Scan(&fl_num)
	n := int(fl_num)

	fmt.Printf("The number is: %d\n", n)
}
