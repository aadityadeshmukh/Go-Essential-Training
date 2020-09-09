package main

import(
	"fmt"
)

func main()  {
	a := 10
	if a > 5{
		fmt.Printf("a is greater\n")
	}	

	if a > 5 && a < 15{
		fmt.Printf("a is just right\n")
	}
	
	x := 11.0
	b := 15.0
	if frac:=x/b; frac > 0.5 {
		fmt.Printf("x is greater than half of b\n")
	}

	y := 2
	switch y{
	case 1:
		fmt.Printf("y is one\n")
	case 2:
		fmt.Printf("y is two\n")
	default:
		fmt.Printf("Dont know what y is\n")
	}

	switch {
	case y > 1:
		fmt.Printf("y is greater than 1\n")
	case y > 10:
		fmt.Printf("y is greater than 10\n")
	default:
		fmt.Printf("Huh?\n")
	}
}	