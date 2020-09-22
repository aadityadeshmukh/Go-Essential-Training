package main 
import(
	"fmt"
)

func add(a int, b int) int{
	return a + b
}

func divmod(a int, b int) (int, int){
	return a/b , a%b
}

func main(){
	val := add(7,10)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("\ndiv is %d and mod is %d" , div, mod)
}