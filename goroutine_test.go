package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i <= 1000000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)

}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Nichola Saputra"
		fmt.Println("Selesai Mengirim Data Ke Channel")
	}()

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Nichola Saputra"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func Menerima(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Nichola Saputra"
}

func Mengambil(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go Menerima(channel)
	go Mengambil(channel)

	time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Nichola"
	channel <- "Saputra"

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan Ke " + strconv.Itoa(i)
		}
		// Pastikan close di channelnya supaya tidak terkena deadlock
		close(channel)
	}()

	// Menerima Data Yang Tidak Jelas Berapa Jumlahnya
	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectMultipleChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	select {
	case data := <-channel1:
		fmt.Println("Data dari channel 1", data)
	case data := <-channel2:
		fmt.Println("Data dari channel 2", data)
	}

	select {
	case data := <-channel1:
		fmt.Println("Data dari channel 1", data)
	case data := <-channel2:
		fmt.Println("Data dari channel 2", data)
	}
}
