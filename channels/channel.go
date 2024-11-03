package channels

import (
	"fmt"
	"time"
)

func ChannelExample() {
	c := make(chan string)

	go func() {
		fmt.Println("Sending Something on the channel")
		c <- "Vansham"
		fmt.Println("Sent something on the channel")
	}()

	time.Sleep(5 * time.Second)
	// fmt.Println(<-c)

	fmt.Println("ending main program")
}

func BufferedChannelExample() {

	ch := make(chan string, 10)

	go func() {
		fmt.Println("Sending Something on the channel")
		for i := 1; i <= 11; i++ {
			ch <- "Vansham"
		}
		// ch <- "Aggarwal"
		fmt.Println("Sent something on the channel")
	}()

	// fmt.Println(<-ch)
	time.Sleep(time.Second * 5)
	fmt.Println("Ending Main")
}
