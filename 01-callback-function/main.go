package main

import (
	"fmt"
	"sync"
	"time"
)

//callback function haha()
func haha() {
	defer wg.Done()
	time.Sleep(time.Second * 5)
	fmt.Println("Hello, my name is haha")
}

//callback function tony()
func tony() {
	defer wg.Done()
	time.Sleep(time.Second * 5)
	fmt.Println("Hello, my name is tony")
}

//args: callback function funcName(with or without arguments){}
func runner(funcName func()) {
	funcName()
	fmt.Println("runner finished")
}

//WaitGroup to make runner()  running background
var wg sync.WaitGroup

func main() {
	fmt.Println("call to haha()")
	wg.Add(1)
	go runner(haha)
	fmt.Println("call to haha()")

	fmt.Println("call to tony()")
	wg.Add(1)
	go runner(tony)
	fmt.Println("call to tony()")
	wg.Wait()
	//runner(haha)
	//runner(tony)


}
