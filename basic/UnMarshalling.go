package main
import 
(
"encoding/json"
"fmt"
)
type MyObject struct {
	Number int `json:"Roll Number"`
	Word string `json:"Name"`
}

func main(){
	jsonBytes := []byte(`{"Roll Number":5, "Name":"Packt"}`)
	var object MyObject
	err := json.Unmarshal(jsonBytes, &object)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Roll Number is %d, Name is %s", object.Number, object.Word)
}
