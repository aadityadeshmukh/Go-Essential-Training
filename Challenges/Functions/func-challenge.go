package main
import(
"fmt"
"net/http"
)

func contentType (url string) (string, error){
	resp, err := http.Get(url)
	
	if err != nil{
		return "Error", fmt.Errorf("Cannot perform a get call: %s \n", err)
	}
	defer resp.Body.Close()
	return resp.Header.Get("Content-Type"), nil
}
func main()  {
	val, err := contentType("https://linkedin.com")
	if err != nil{
		fmt.Printf("ERROR: %s\n", err)
	}else{
		fmt.Printf(val)
	}
}