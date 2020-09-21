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
	stocks["TSLA"] = 420.69
	fmt.Println(stocks["TSLA"])
	delete(stocks, "MSFT")
	fmt.Println(stocks)

	value, ok := stocks["GOOG"]
	if !ok {
		fmt.Println("Not found")
	}else{
		fmt.Println(value)
	}
}