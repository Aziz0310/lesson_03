package main

import "fmt"

func main() {

	var i, j, k, row int

	fmt.Print("Enter Hollow Rhombus Star Pattern Rows = ")
	fmt.Scanln(&row)

	fmt.Println("Hollow Rhombus Star Pattern")

	for i = 1; i <= row; i++ {
		for j = 1; j <= row-i; j++ {
			fmt.Printf(" ")
		}
		for k = 1; k <= row; k++ {
			if i == 1 || i == row || k == 1 || k == row {
				fmt.Printf("* ")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
}
