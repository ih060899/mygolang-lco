package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

func main() {
	//go greeting("Hello")
	//greeting("world")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)
}

//func greeting(message string) {
//	for i := 0; i < 6; i++ {
//		//time.Sleep(3 * time.Millisecond)
//		fmt.Println(message)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}