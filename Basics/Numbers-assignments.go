package main
import(
	"fmt"
)
func main()  {
	
	// var a float64 //type assignment comes after the declaration
	// var b float64 //if no value is assigned it is defaulted as 0

	// a:=1 //assignment
	// b:=2

	// //define and assignment
	// a:=1.0
	// b:=2.0

	//multiline define and assingment
	a, b := 1.0, 2.0
	
	fmt.Printf("a=%v, is of type %T\n", a, a) //%v and %T are verbs
	fmt.Printf("b=%v, is of type %T\n", b, b) //%v prints the value and %T prints the type

	var mean float64
	mean=(a+b)/2.0
	fmt.Printf("mean=%v, is of type %T\n", mean, mean)
}