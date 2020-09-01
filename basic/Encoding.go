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
	object := MyObject{5, "Packt"}
	oJson, _ := json.Marshal(object)
	fmt.Printf("%s\n", oJson)
}
