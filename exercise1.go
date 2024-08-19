package main
// Exercise: Loops and Functions
// https://go.dev/tour/flowcontrol/8
import (
	"fmt"
	"math"
)

const LIMIT = 1e-15

func Sqrt(x float64) float64 {
	loops := 1
	z := float64(x)
	for {
		za := z - (z*z-x)/(2*z)
		fmt.Println(za)
		if math.Abs(z-za) < LIMIT {
			fmt.Print("Loop nº", loops, " = ")
			return z
		} else {
			z = za
		}
		loops++
	}
}

func main() {
	test := float64(1)
	fmt.Println(Sqrt(test))
	fmt.Println(math.Sqrt(test))
}