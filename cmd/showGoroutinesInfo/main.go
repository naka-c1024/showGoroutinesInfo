package main

import (
	"sync"

	"github.com/naka-c1024/showGoroutinesInfo"
)

func main() {
	showGoroutinesInfo.Do("before main")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineFirst()
	}()

	wg.Wait()
	showGoroutinesInfo.Do("after main")
}

func goroutineFirst() {
	showGoroutinesInfo.Do("goroutineFirst")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineSecond()
	}()

	wg.Wait()
}

func goroutineSecond() {
	showGoroutinesInfo.Do("goroutineSecond")
}
