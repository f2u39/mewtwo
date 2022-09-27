package main

import (
	"fmt"
	"time"
)

func main() {
	s := "일초 일분 한시간"
	p := &s

	go func() {
		for {
			fmt.Println("goroutine accessing to the pointer", *p)
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		fmt.Println("main() accessing to the pointer", *p)
		time.Sleep(1 * time.Second)
	}
}
