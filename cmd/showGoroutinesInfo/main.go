package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
)

func main() {
	showGoroutinesInfo("before main")
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineFirst()
	}()

	wg.Wait()
	showGoroutinesInfo("after main")
}

func goroutineFirst() {
	showGoroutinesInfo("goroutineFirst")

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineSecond()
	}()

	wg.Wait()
}

func goroutineSecond() {
	showGoroutinesInfo("goroutineSecond")
}

// =======================================================

type GoroutineInfo struct {
	ID      string
	State   string
	Created string
}

var m sync.Mutex

func showGoroutinesInfo(regionName string) {
	m.Lock()
	defer m.Unlock()

	fmt.Printf("\n=== goroutines info: %s ===\n", regionName)

	buf := make([]byte, 1024)
	n := runtime.Stack(buf, true)
	str := string(buf[:n])

	lines := strings.Split(str, "\n")

	var goroutines []GoroutineInfo
	var goroutine GoroutineInfo
	for _, line := range lines {
		if strings.HasPrefix(line, "goroutine") {
			if goroutine.ID != "" {
				goroutines = append(goroutines, goroutine)
				goroutine = GoroutineInfo{}
			}
			lineSplit := strings.Split(strings.TrimSpace(line), " ")
			if len(lineSplit) >= 3 {
				goroutine.ID = lineSplit[1]
				state := strings.Join(lineSplit[2:], " ")
				goroutine.State = strings.Trim(strings.Trim(strings.Trim(state, ":"), "["), "]")
			}
		} else if strings.HasPrefix(line, "created by") {
			goroutine.Created = strings.TrimSpace(strings.TrimPrefix(line, "created by"))
		}
	}
	if goroutine.ID != "" {
		goroutines = append(goroutines, goroutine)
	}

	fmt.Printf("\nnum goroutines -> %d\n\n", len(goroutines))

	for _, g := range goroutines {
		fmt.Printf("Goroutine ID: %s\nState: %s\n", g.ID, g.State)
		if g.Created != "" {
			fmt.Printf("Created at: %s\n", g.Created)
		}
		fmt.Printf("\n")
	}
}
