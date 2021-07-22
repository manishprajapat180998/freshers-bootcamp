package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var sum = int(0)

func rating(j int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum += rand.Intn(20)
	n := rand.Intn(5)
	time.Sleep(time.Duration(n) * time.Millisecond)
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go rating(i, &wg)
	}
	wg.Wait()
	x := sum / 200
	fmt.Println(x)
}
