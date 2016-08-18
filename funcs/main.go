package main

import "fmt"
import "math"

func pow(k float64) func(n float64) float64 {
	f := func(n float64) float64 {
		return math.Pow(n, k)
	}
	return f
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func sum(array []float64) float64 {
	result := 0.0
	for _, v := range array {
		result += v
	}
	return result
}

func half(n int) (int, bool) {
	r1 := n / 2
	r2 := n % 2
	var r22 bool
	if r2 == 0 {
		r22 = true
	} else {
		r22 = false
	}
	return r1, r22
}

func main() {
	//xs := []float64{98, 93, 77, 82, 83}
	//fmt.Println(average(xs))
	arr := []float64{2.2, 3.3, 4.7}
	fmt.Println(sum(arr))
	square := pow(2)
	fmt.Println(square(4))

	fmt.Println(half(5))

}
