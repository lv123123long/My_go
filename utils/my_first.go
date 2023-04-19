package main

import (
"fmt"
"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				fmt.Println("会姐，牛，牛，牛")
			}
		}()
	}

	wg.Wait()
}