package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0;
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2*z)
		fmt.Printf("第%v次：z=%v\n", i+1, z)
	}
    return z
}

func main() {
	fmt.Println(Sqrt(2))
}

