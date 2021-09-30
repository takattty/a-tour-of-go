package main

import "fmt"

func main() {
	sum := 1
	// for i := 0; i < 1000; i++ {
	// 	fmt.Println(sum)
	// 	sum += i
	// }
	for sum < 1000000 {
		fmt.Println(sum)
		sum += sum
	}
	fmt.Println(sum)
}
