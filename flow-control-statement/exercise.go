package main

import (
	"fmt"
	"math"
)

const D = 0.0000001

func Sqrt(x float64) float64 {
	z := float64(1)
	d := 0.0
	// for i := 0; i < 10; i++ {
	// 	d = z
	// 	z -= (z*z - x) / (2 * z)
	// 	// z = z - (z*z-x)/(2*z)
	// 	if math.Abs(z-d) < D {
	// 		break
	// 	}
	// }
	for math.Abs(z-d) > D {
		d = z
		z -= (z*z - x) / (2 * z)
		// d, z = z, z - (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
