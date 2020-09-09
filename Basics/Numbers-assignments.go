package main
import(
	"fmt"
)
func main()  {
	var a int //type assignment comes after the declaration
	var b int //if no value is assigned it is defaulted as 0

	a=1 //assignment
	b=2

	fmt.Printf("a=%v, is of type %T\n", a, a) //%v and %T are verbs
	fmt.Printf("b=%v, is of type %T\n", b, b) //%v prints the value and %T prints the type

	var mean int
	mean=(a+b)/2
	fmt.Printf("mean=%v, is of type %T\n", mean, mean)
}