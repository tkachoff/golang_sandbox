package main

import "fmt"

func main() {
	/*	fmt.Print("Enter a number: ")
		var input float64
		fmt.Scanf("%f", &input)

		output := input * 2

		fmt.Println(output)
	*/
	i := 1
	for i <= 10 {
		fmt.Print(i)
		if i%2 == 0 {
			fmt.Println("\teven")
		} else {
			fmt.Println("\todd")
		}
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

}
