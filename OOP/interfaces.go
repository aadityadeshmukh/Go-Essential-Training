package main 
import (
	"fmt"
	"math"
)
type Square struct{
	Length float64
}
func (sq *Square) Area() float64{
	return sq.Length * sq.Length
}

type Circle struct{
	Radius float64
}

func (ci *Circle) Area() float64{
	return ci.Radius * ci.Radius * math.Pi
}

type Shape interface{
	Area() float64
}

func sumShapes(shapes []Shape) float64{
	total := 0.0
	for _, shape := range shapes{
		total += shape.Area()
	}

	return total
}

func main(){
	sqr := Square{30}
	fmt.Println(sqr.Area())

	cir := Circle{30}
	fmt.Println(cir.Area())

	grpShapes := []Shape{&sqr, &cir}
	sumAreas := sumShapes(grpShapes)
	fmt.Println(sumAreas)
}