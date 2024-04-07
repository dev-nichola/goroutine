package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyncronus(group *sync.WaitGroup, number int) {
	defer group.Done()
	group.Add(1)

	fmt.Println("Hello World + ", number)
	time.Sleep(5 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsyncronus(group, i)
	}

	group.Wait()
	fmt.Println("Complete")
}
