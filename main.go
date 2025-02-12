package main

import (
	"fmt"
	"math/rand/v2"
)

func average_num(n int) float64 {
	var sum float64
	var count float64
	var roll float64
	var rand_num int

	for i:=0; i<n; i++ {
		rand_num = rand.IntN(20) + 1
		roll = float64(rand_num)

		sum += roll
		count += 1.0
	}

	return sum/count
}

func main() {
	var input_num int

//	a_num := rand.IntN(20)
//	fmt.Println(a_num)
//	fmt.Println(float32(a_num))
//	a_num := rand.IntN(20)
//	fmt.Println(a_num)
//	fmt.Println(float32(a_num))
//	a_num := rand.IntN(20)
//	fmt.Println(a_num)
//	fmt.Println(float32(a_num))
//	a_num := rand.IntN(20)
//	fmt.Println(a_num)
//	fmt.Println(float32(a_num))

	fmt.Print("Input number of rolls: ")

	fmt.Scanln(&input_num)

	average := average_num(input_num)

	fmt.Printf("Average result: %f", average)
}
