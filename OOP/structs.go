package main 
import (
	"fmt"
	"os"
)
type Trade struct{
	Symbol string
	Volume int
	Price float64
	Buy bool
}

func NewTrade(symbol string, volume int, price float64, buy bool) (*Trade, error){
	if symbol == ""{
		return nil, fmt.Errorf("symbol can't be empty")
	}
	if volume <= 0{
		return nil, fmt.Errorf("volume cannot be 0")
	}
	if price <= 0{
		return nil, fmt.Errorf("price cannot be negative")
	}

	trade := Trade{symbol, volume, price, buy}
	return &trade, nil
}

func (t *Trade) Value() float64{
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}
func main(){

	t := Trade{"MSFT", 10, 300.64, true}
	fmt.Println(t)
	fmt.Printf("%+v\n", t) //verbose print

	t1 := Trade{Symbol: "APPL", Volume: 30, Price: 2000.3, Buy: false}
	fmt.Println(t1)

	t2 := Trade{}
	fmt.Println(t2)

	t3 := Trade{"TSLA", 20, 350.64, true}
	fmt.Println(t3.Value())

	t4, err := NewTrade("GOOG", 15, -1500.0, true)

	if err != nil{
		fmt.Printf("error")
		os.Exit(1)
	}

	fmt.Println(t4.Value())
}