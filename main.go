package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	channel := make(chan int)
	for i := 1; i <= 10; i++ {
		fmt.Printf("Pemain %v Memulai Permainan\n", i)
		go process(i, channel)
	}
	fmt.Printf("Pemain %v Kalah ~", <-channel)
}

func process(player int, channel chan int) {
	for {
		time.Sleep(1 * time.Second)
		a := rand.NewSource(time.Now().UnixNano())
		random := rand.New(a)
		d := random.Intn(100) % 10
		fmt.Printf("Pemain %v Memegang Petasan\n", player)
		if d == 0 {
			channel <- player
			return
		}
	}
}
