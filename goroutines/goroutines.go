package goroutines

import (
	"fmt"
	"time"
)

func ChannelExample() {

	messages := make(chan string);

	go func() { 
		time.Sleep(time.Second * 5)
		messages <- "helluu"}()

	msg := <-messages

	fmt.Println(msg)
}