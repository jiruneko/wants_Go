package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	x := complex(2.5, 3.1)
	y := complex(10.2, 2)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)
	fmt.Println(x / y)
	fmt.Println(real(x))
	fmt.Println(imag(x))
	fmt.Println(cmplx.Abs(x))

	var z int = 10
	var o float64 = 30.2
	var sum1 float64 = float64(z) + o
	var sum2 int = z + int(o)
	fmt.Println(sum1, sum2)

	var Ω int = 10
	var b byte = 100
	var sum3 int = Ω + int(b)
	var sum4 byte = byte(Ω) + b
	fmt.Println(sum3, sum4)
}
