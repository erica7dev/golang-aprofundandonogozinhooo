package main 

 import (
	"fmt"
	"runtime"
	"sync"
 )

 func main(){
		fmt.Println("CPU: ", runtime.NumCPU())
		fmt.Println("Goroutines: ", runtime.NumGoroutine())

		count := 0
		totalGoroutines := 100

		var wg sync.WaitGroup
		wg.Add(totalGoroutines)

		for i := 0; i < totalGoroutines; i++ {
			go func() {
				v := count
				runtime.Gosched()
				v++
				count = v
				wg.Done()
			}()
			fmt.Println("Goroutines:", runtime.NumGoroutine())
		}

		wg.Wait()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		fmt.Println("Valor final:", count)
 }