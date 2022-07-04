package main

import (
	"runtime"
	"sync"
	"fmt"
	"sync/atomic"
)

func main(){
	fmt.Println("CPU: ", runtime.NumCPU())

	var contador int64

	contador = 0
	totalGoroutines := 100

	var wg sync.WaitGroup
	wg.Add(totalGoroutines)

	for i := 0; i < totalGoroutines; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched()
			fmt.Println("contador: \t", atomic.LoadInt64(&contador))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Valor final: \t", contador)
}