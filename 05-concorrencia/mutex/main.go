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
		totalGoroutines := 10

		var wg sync.WaitGroup
		wg.Add(totalGoroutines)

		var mu sync.Mutex

		for i := 0; i < totalGoroutines; i++ {
			go func() {
				mu.Lock()
				v := count
				runtime.Gosched()
				v++
				count = v
				mu.Unlock()
				wg.Done()
			}()
			fmt.Println("Goroutines:", runtime.NumGoroutine())
		}

		wg.Wait()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		fmt.Println("Valor final:", count)
 }