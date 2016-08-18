package main

import "fmt"

func main() {
	a := []float64{
		98,
		93,
		77,
		82,
		83,
	}

	b := []float64{
		97,
		38,
	}

	slice2 := append(a, b...)

	total := 0.0
	//	for i := 0; i < len(a); i++ {
	//		total += a[i]
	//	}
	for _, value := range slice2 {
		total += value
	}
	fmt.Println(slice2)
	fmt.Println(b)
	fmt.Println(total / float64(len(slice2)))

	dict := make(map[string]int)
	dict["artsiom"] = 29
	dict["maryna"] = 25

	age, ok := dict["max"]
	fmt.Println(age, ok)

	age2, ok2 := dict["artsiom"]
	fmt.Println(age2, ok2)
}
