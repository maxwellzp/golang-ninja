package main

import (
	"fmt"
	"time"
)

func main() {
	//// #1
	// var msg chan string
	// fmt.Println(msg)

	// msg = make(chan string)
	// fmt.Println(msg)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	// Write to channel
	// 	msg <- "My channel"
	// }()

	// value := <-msg // Reading from channel and writing to channel is a blocking operation
	// fmt.Println(value)

	//// #2

	// msg := make(chan string, 3)

	// // Небуфирезированный канал блокируется на запись
	// // Буферизированные каналы дают возможность записывать значения, без блокировки
	// msg <- "My channel 1"
	// msg <- "My channel 2"
	// msg <- "My channel 3"

	// // Reading from channel and writing to channel is a blocking operation
	// fmt.Println(<-msg)
	// fmt.Println(<-msg)
	// fmt.Println(<-msg)

	//// #3
	// msg := make(chan string, 3)

	// msg <- "My channel 1"
	// msg <- "My channel 2"
	// msg <- "My channel 3"
	// close(msg)

	// for m := range msg {
	// 	fmt.Println(m)
	// }

	//// #4
	// msg := make(chan string, 3)

	// msg <- "My channel 1"
	// msg <- "My channel 2"
	// msg <- "My channel 3"

	// close(msg)

	// for {
	// 	value, ok := <-msg
	// 	if !ok {
	// 		fmt.Println("channel closed")
	// 		break
	// 	}

	// 	fmt.Println(value)
	// }

	//// #5
	message1 := make(chan string)
	message2 := make(chan string)

	go func() {
		for {
			message1 <- "Канал 1. Прошло 200 мс."
			time.Sleep(time.Millisecond * 200)
		}
	}()

	go func() {
		for {
			message2 <- "Канал 2. Прошла 1 c."
			time.Sleep(time.Second)
		}
	}()

	for {
		// Select позволяет нам неблокировать (чтение) один канал другим каналом.
		select {
		case msg := <-message1:
			fmt.Println(msg)
		case msg := <-message2:
			fmt.Println(msg)
		default:

		}
	}
}
