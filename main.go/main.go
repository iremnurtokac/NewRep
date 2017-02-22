package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	str1 := "<>"
	str2 := "<>"
	for i := 0; i < 2; i++ {
		str1 = strings.Replace(str1, "<>", "<<>>", -1)
		str2 = strings.Replace(str2, "<>", "<<>>", -1)
		str2 = strings.Replace(str2, "<>", "<> <>", -1)
	}

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- str1

	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- str2
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
