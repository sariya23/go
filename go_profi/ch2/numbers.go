package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Repeat("#", 10) + "COMPLEX" + strings.Repeat("#", 10))
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)

	var c3 complex64 = complex64(c1 + c2)
	fmt.Println("c3:", c3)
	fmt.Printf("Type of c3: %T\n", c3)
	cZero := c3 - c3
	fmt.Printf("Type of cZero is %T\n", cZero)
	fmt.Println(cZero)
	fmt.Println()

	fmt.Println(strings.Repeat("#", 10) + "NUMBERS" + strings.Repeat("#", 10))
	x := 11
	k := 5
	intdiv := 11 / 5
	fmt.Println(intdiv)
	fmt.Printf("Type of intdiv is %T\n", intdiv)

	div := x / k
	fmt.Println("div:qwqw ", div)

	var m, n float64
	m = 1.223
	fmt.Println("m, n - ", m, n)
	y := 4.0 / 2
	fmt.Printf("Type of y is %T\n", y)

	divFloat := float64(x) / float64(k)
	fmt.Printf("Type of divFloat is %T", divFloat)
}
