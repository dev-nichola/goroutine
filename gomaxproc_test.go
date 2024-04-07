package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU : ", totalCpu)

	// Mengubah Jumlah Thread
	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Go Routine : ", totalGoroutine)

	group.Wait()
}
