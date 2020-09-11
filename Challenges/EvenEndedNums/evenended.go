// Even ended numbers like 11, 121 resulting from product of 4 digit numbers
package main 
import (
	"fmt"
)
func main(){
	// multiply every 4 digit number with every other 4 digit number
	//convert each product to string
	//check if 1st and last strings are same
	a:=0
	for i:= 1000; i <= 9999; i++{
		for j:=i; j <= 9999; j++{
			prod := i*j
			strProd := fmt.Sprintf("%d", prod)
			length := len(strProd)
			if strProd[0]  == strProd[length-1] {
				// fmt.Println(strProd)
				a++
			}
		}
	}
	fmt.Println(a)
}