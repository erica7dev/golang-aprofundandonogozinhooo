package main

import (
	"fmt"
)

var x int = 40
var y bool = false
var z string = "Menina"

func main(){
	juntaTudo := fmt.Sprintf("%v\n%v\n%v\n", x,y,z)
	fmt.Println(juntaTudo)
}