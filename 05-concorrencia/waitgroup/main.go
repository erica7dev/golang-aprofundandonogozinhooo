package main

import (
	"fmt"
	"runtime"
	"math"
)

func main(){
	threads := 50

	fmt.Println("Inicio. Goroutines: ", runtime.NumGoroutine())

	for i:=0; i < threads; i++ {
		k := i
		go func() {
			fmt.Print("Thread #", k, "inicializando...") //marca o indice da thread
			for j:=0;j<10;j++ {
				fmt.Print("Thread #", k,"contagem: ", j, "Operação: ", math.Sqrt(float64(j)*math.Pi), ".\n") //10 operações de arrendondar raiz de pi
			}
			threads-- //threads finalizadas
			fmt.Print("Thread #", k, "finalizada. Restam ", threads, "threads.\n")
		}()
	}

	fmt.Println("Meio. Goroutines: ", runtime.NumGoroutine())
	for threads != 0 {

	}

	fmt.Println("Fim. Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Fim do programa.")
}