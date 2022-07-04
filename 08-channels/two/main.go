package main 

import (
	"fmt"
	"sync"
)

//recebendo info do channel par e impar no converge
func main(){
	par := make(chan int)
	impar := make(chan int)
	converge := make(chan int)

	go envia(par,impar)
	go recebe(par,impar,converge)

	for v:= range converge{
		fmt.Println("Valor recebido:",v)
	}
}

func envia(p, i chan int){
	x := 100
	for n := 0; n < x; n++{
		if n%2 == 0{
			//enviando info
			p <- n 
		}else{
			i <- n
		}
	}
	close(p)
	close(i)
}

func recebe(p, i, c chan int){
	var wg sync.WaitGroup
	wg.Add(1)
	go func(){
		for v := range p{
			//channel recebe
			c <- v
		}
		wg.Done()
	}()
	wg.Add(1)
	go func(){
		for v := range i{
			//channel recebe
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}