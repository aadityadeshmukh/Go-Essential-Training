package main
import(

"fmt"

)
func main()  {
	numset := []int {3, 2, 84, 66, 128, 5}
	largestNum := numset[0]
	for i:=0; i < len(numset); i++{
		if numset[i] > largestNum {
			largestNum = numset[i]
			fmt.Printf("New max: %v\n", largestNum)
		}
	}
	fmt.Printf("Largest number: %v\n" , largestNum)
}