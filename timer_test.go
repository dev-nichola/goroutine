package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	channel := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <-channel.C
	fmt.Println(time)
}
