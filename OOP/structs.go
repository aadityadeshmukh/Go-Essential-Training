package main 
import (
	"fmt"
)
type Trade struct{
	Symbol string
	Volume int
	Price float64
	Buy bool
}
func main(){

	t := Trade{"MSFT", 10, 300.64, true}
	fmt.Println(t)

}