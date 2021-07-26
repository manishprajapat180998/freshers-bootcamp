package main

import (
	"fmt"
	"sync"
)
var balance = int(500)

func deposit(wg *sync.WaitGroup,mutex *sync.Mutex, value int){
	defer wg.Done()
	mutex.Lock()
	balance+=value
	fmt.Println(balance)
	mutex.Unlock()
}
func withdraw(wg *sync.WaitGroup,mutex *sync.Mutex, val int){
	defer wg.Done()
	mutex.Lock()
	if balance<val{
		fmt.Println("cannot withdraw")
	}else{
		balance-=val
		fmt.Println(balance)
	}
	mutex.Unlock()
}
func main(){

	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	wg.Add(2)
	var value = int(1000)
	var val = int(500)
	go deposit(&wg,mutex,value)
	go withdraw(&wg,mutex,val)
	wg.Wait()

}
