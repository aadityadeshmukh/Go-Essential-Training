package main 
import(
	"fmt"
	"math"
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

//Error handling
func sqrt( inp float64) (float64, error){
	if inp < 0 {
		return 0.0, fmt.Errorf("Cannot square root a negative number.\n")
	}

	return math.Sqrt(inp), nil
}

func cleanup (inp string){
	fmt.Printf("Cleaning up %s\n", inp)
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

	result, err := sqrt(3.0)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}

	result, err = sqrt(-2.0)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}

	defer cleanup("A")
	defer cleanup("B")

	fmt.Println("Worker")
}