package main

import (
	"fmt"
	"runtime"
)

//criando types
type num int
var x num

var y int

func main(){

	x = 50
	fmt.Println(x)
	//converter y em num
	y = int(x)
	fmt.Printf("%T\n", y)
	fmt.Print(runtime.GOARCH)
	fmt.Print(runtime.GOOS)
}