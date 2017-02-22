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
	for i := 0; i < 5; i++ {
		str1 = strings.Replace(str1, "<>", "<<>>", -1)
		str2 = strings.Replace(str2, "<>", "<<>>", -1)
		str2 = strings.Replace(str2, "<>", "<> <>", -1)

		go func() {
			temp := str1
			time.Sleep(time.Second * time.Duration(i))
			c1 <- temp

		}()
		go func() {

			temp := str2
			time.Sleep(time.Second * time.Duration(i))
			c2 <- string(temp)

		}()
	}

	for i := 0; i < 10; i++ {
		select {
		case msg1 := <-c1:

			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}

	}

	msg1 := <-c1
	fmt.Println(msg1)
	msg2 := <-c2
	fmt.Println(msg2)

}
