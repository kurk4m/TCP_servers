package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	// #1
	// fmt.Println("Scanner test")
	// address := "scanme.nmap.org:80"
	// _, err := net.Dial("tcp", address)
	// if err != nil {
	// 	fmt.Print("error")
	// }
	// fmt.Print("port opened!")

	// #2 range
	// for i := 1; i <= 1024; i++ {
	// 	address := fmt.Sprintf("scanme.nmap.org:%d", i)
	// 	conn, err := net.Dial("tcp", address)
	// 	if err != nil {
	// 		continue
	// 		// fmt.Print("error")
	// 	}
	// 	conn.Close()
	// 	fmt.Printf("port opened! %d\n", i)
	// }

	// #3 concurrent
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)

			if err != nil {
				fmt.Printf("Error connecting to %s: %v\n", address, err)
				return // return to exit the goroutine if there's an error
			}

			conn.Close()
			fmt.Printf("port opened! %d\n", j)
		}(i)
	}
	wg.Wait()
}

func handleScanner() {}

func handleScannerRange() {}
