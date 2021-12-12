package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("CPUs:", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

		var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func(){
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
	//Gosched yields the processor, allowing other goroutines to run.
}