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
	fmt.Printf("%+v\n", t) //verbose print

	t1 := Trade{Symbol: "APPL", Volume: 30, Price: 2000.3, Buy: false}
	fmt.Println(t1)
}