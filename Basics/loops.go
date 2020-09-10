package main
import (
	"fmt"
)
func main(){
	for i:=0; i < 3; i++{
		if i > 1{
			break
		}
		fmt.Printf("i is %v\n", i)
	}

	fmt.Println("--------------")
	for i:=0; i < 3; i++{
		if i < 1{
			continue
		}
		fmt.Printf("i is %v\n", i)
	}

	fmt.Println("-----------")
	a:=0
	for a < 3 { //similar to while loop in other languages
		fmt.Printf("a is %v\n", a)
		a++
	}

	fmt.Println("-----------")
	b:=0
	for { //similar to while loop in other languages
		if(b > 4){
			break
		}
		fmt.Printf("a is %v\n", a)
		b++
	}

}