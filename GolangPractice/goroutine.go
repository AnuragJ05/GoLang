package main

import (
	"fmt"
	"time"
)

func main() {
	var c1 = make(chan string)
	var c2 = make(chan string)
	go count("Anurag", c1)
	go count("Jain", c2)
	for {
		select {
		case msg, open := <-c1:
			if !open {
				break
			}
			fmt.Println(msg)

		case msg, open := <-c2:
			if !open {
				break
			}
			fmt.Println(msg)
		}
	}
}

func count(value string, c chan string) {
	for i := 1; i <= 10; i++ {
		c <- value
		time.Sleep(1)
	}
	close(c)
}
