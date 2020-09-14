package main 
import (
	"fmt"
)
func main(){
	looney := []string {"bugs", "daffy", "traz"}
	fmt.Printf("Looneys = %v of the type %T\n" , looney, looney)
	fmt.Println(len(looney))
	fmt.Println(looney[1:])

	for i:= 0; i< len(looney); i++{
		fmt.Println(looney[i])
	}

	for i:= range looney{
		fmt.Println(i)
	}

	for i, name:= range looney{
		fmt.Printf(name)
		fmt.Printf(" ")
		fmt.Println(i)
	}

	for _, name:= range looney{
		fmt.Printf(name)
		fmt.Printf(" \n")
	}

	looney = append(looney, "elmer")
	fmt.Printf("Looneys = %v of the type %T\n" , looney, looney)
}