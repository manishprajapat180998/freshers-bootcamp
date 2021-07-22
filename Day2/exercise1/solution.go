package main

import (
	"fmt"
)

var m = make(map[string]int)

func freq_count(strin_channel <-chan string, strin_channel1 chan<- string) {
	str_temp := <-strin_channel
	x := len(str_temp)
	for j := 0; j < x; j++ {
		s := string(str_temp[j])
		m[s]++
	}
	strin_channel1 <- str_temp
}

func main() {
	stri := make([]string, 5)
	stri[0] = "quick"
	stri[1] = "brown"
	stri[2] = "fox"
	stri[3] = "lazy"
	stri[4] = "dog"

	strin_channel := make(chan string, len(stri))
	strin_channel1 := make(chan string, len(stri))

	for i := 0; i < 5; i++ {
		go freq_count(strin_channel, strin_channel1)
	}
	for i := 0; i < 5; i++ {
		strin_channel <- stri[i]
	}
	for i := 0; i < 5; i++ {
		<-strin_channel1
	}

	for k, _ := range m {
		fmt.Println(k, m[k])
	}
}
