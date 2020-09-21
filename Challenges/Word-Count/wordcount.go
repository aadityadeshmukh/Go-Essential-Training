package main 
import(
	"fmt"
	"strings"
)
func main(){
	text := `
	Needles and pins
	Needles and pins
	Sew me a sail
	To catch me the wind
	`
	//create a slice of the text
	words := strings.Fields(text) //ignores white space

	//Create an empty map
	counts := map[string]int{}

	//iterate over the words slice
	for _, word := range words{
		counts[strings.ToLower(word)]++
	}

	fmt.Println(counts)
}