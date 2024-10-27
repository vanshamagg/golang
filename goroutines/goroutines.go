package goroutines

import (
	"fmt"
	"time"
)

func ChannelExample() {

	messages := make(chan string)

	go func() {
		time.Sleep(time.Second * 5)
		messages <- "helluu"
	}()

	msg := <-messages

	fmt.Println(msg)
}

func SelectExample() {

	c1 := make(chan string)
	c2 := make(chan string)

	// goroutine 1

	go func() {
		time.Sleep(time.Second * 5)
		c1 <- "done"
	}()

	// go routine 2
	go func() {
		time.Sleep(time.Second * 5)
		c1 <- "done as well"
	}()

	for i := 0; i < 2; i++ {
		select {
		case ms1 := <-c1:
			fmt.Println("received  ", ms1)
		case ms2 := <-c2:
			fmt.Println("received ", ms2)
		}
	}
}

func BufferedChannelsExample() {

	bufch := make(chan string, 4)

	go func() {
		bufch <- "hello"
		bufch <- "there"
		bufch <- "whats up"
		bufch <- "you good?"
	}()

	output, ok := <-bufch
	fmt.Println(output)
	fmt.Println(ok)
	fmt.Println(<-bufch)
}
