package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Millisecond)
}

func main() {
	runtime.GOMAXPROCS(20)
	MAX_TASKS := 10000
	var wg sync.WaitGroup

	wg.Add(MAX_TASKS)
	start := time.Now()

	for i := 0; i < MAX_TASKS; i++ {
		go worker(&wg)
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
