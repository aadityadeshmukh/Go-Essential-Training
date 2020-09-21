package main 
import (
	"fmt"
)
func main(){

	stocks := map [string] float64{
		"AMZN" : 100.1,
		"GOOG" : 1000.2,
		"MSFT" : 300.8,
	}

	fmt.Println(stocks)
	fmt.Println(len (stocks))
	fmt.Println(stocks["AMZN"])
	fmt.Println(stocks["TSLA"])

	value, ok := stocks["GOOG"]
	if !ok {
		fmt.Println("Not found")
	}else{
		fmt.Println(value)
	}
}