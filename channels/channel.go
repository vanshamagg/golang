package channels

import (
	"fmt"
	"time"
)

func ChannelExample() {
	c := make(chan string)

	go func (){
		fmt.Println("Sending Something on the channel")
		c <- "Vansham"
		fmt.Println("Sent something on the channel")
	}()

	time.Sleep(5 * time.Second)
	// fmt.Println(<-c)

	fmt.Println("ending main program")
}