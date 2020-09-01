package main

import (
 "encoding/json"
 "fmt"
)

func main() {
	packt := "marshal"
	jsonPackt, _ := json.Marshal(packt)
	fmt.Println(string(jsonPackt))
}
