package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	<-timer.C

	fmt.Println("timer 1 expired")

	timer2 := time.NewTimer(1 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer 2 expired")
	}()

	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer2 stopped")
	}
}
