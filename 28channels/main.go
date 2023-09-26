package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in golang")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	//fmt.Println(<-myChannel)
	//myChannel <- 5
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		//close(myChannel)
		val, isChannelOpen := <-myChannel
		fmt.Println(isChannelOpen)
		fmt.Println(val)

		//fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5
		close(myChannel)
		//myChannel <- 6
		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
