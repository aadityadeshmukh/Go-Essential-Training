// Fizzbuzz challenge
package main 
import(
	"fmt"
	"math"
)

func main(){
	for i:=1.0; i <= 20.0; i++{
		mod3 := math.Mod(i,3.0)
		mod5 := math.Mod(i,5.0)
		
		if mod3 == 0 && mod5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if mod5 == 0{
			fmt.Println("Buzz")
			continue
		}
		if mod3 == 0{
			fmt.Println("Fizz")
			continue
		}

		fmt.Println(i)

	}
}