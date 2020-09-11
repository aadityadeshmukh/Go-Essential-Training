package main 
import(
	"fmt"
)
func main(){
	book := "Fingerprints of the Gods"
	fmt.Println (book)
	// Size
	fmt.Println (len(book))
	fmt.Printf ("book[0] == %v and is of type %T\n" , book[0] , book[0])

	// Concatenation and slice
	fmt.Println( "The" + book)
	fmt.Println(book[4:11])
	fmt.Println( book[4:] )
	fmt.Println( book[:4] )

	// special stuff

	fmt.Println ("The book is ☺")
	str:= `.......
	I am a poet
	And I didnt even know I was rhyming those words
	......
	`
	println(str)
}