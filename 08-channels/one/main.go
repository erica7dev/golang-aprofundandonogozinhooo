package main

import (
	"fmt"
	"time"
)

func main(){
	//channels declarados e inicializados ao mesmo tempo
	c := make(chan int)
	q := make(chan bool)

	go boop(c,q)

	for{
		select{
			//channel está lendo
		case v, ok := <- c:
			if ok{
				fmt.Println(v)
			}
		case <- q:
			return
		}
	}
}

func boop(c chan <- int, q chan <- bool) {
	for i := 0; i < 10; i++ {
		//escrevendo
		c <- i
	}
	close(c)
	time.Sleep(5 * time.Microsecond)
	//escrevendo
	//fica bloqueada até que outra leia
	q <- true
	close(q)
}