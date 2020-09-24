package main
import(

"fmt"

)
type Point struct{
	X int
	Y int
}

func (pt *Point) Move(dx int, dy int){
	pt.X += dx
	pt.Y += dy
}

type Square struct{
	Center Point
	Length int
}
func (sq *Square) Move(dx int, dy int){
	sq.Center.Move(dx,dy)
}

func (sq *Square) Area() int{
	area := sq.Length * sq.Length
	return area
}

func NewSquare (cen Point, len int) *Square {
	sqr := Square {cen, len}
	return &sqr
}
func main()  {

	center := Point{30, 30}
	fmt.Println(center)

	sqr := NewSquare(center, 40)
	fmt.Println(sqr)

	fmt.Println(sqr.Area())
	sqr.Move(20, 25)
	fmt.Println(sqr)
}

