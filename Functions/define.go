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

func doubleAt(a[]int, i int) {
	a[i] *= 2
}

func double(i *int){
	*i *= 2
}

func main(){
	val := add(7,10)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div is %d and mod is %d\n" , div, mod)

	values := [] int {1, 2, 4, 10}

	doubleAt(values, 2)
	fmt.Println(values)

	num := 4
	double(&num)
	fmt.Println(num)

}