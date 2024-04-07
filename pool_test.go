package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// Membuat Default Data Pool Tinggal Override
	pool := sync.Pool{
		New: func() any {
			return "Default Value"
		},
	}
	// group := sync.WaitGroup{}

	pool.Put("Niko")
	pool.Put("Nichola")
	pool.Put("Saputra")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	// group.Wait()
	time.Sleep(10 * time.Second)
	fmt.Println("Complete")
}
