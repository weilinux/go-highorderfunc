package main

import (
	"fmt"
	"time"
)

//callback function haha()
func haha() {
	time.Sleep(time.Second * 5)
	fmt.Println("Hello, my name is haha")
	//wg.Done()
}

//callback function tony()
func tony() {
	time.Sleep(time.Second * 5)
	fmt.Println("Hello, my name is tony")
	//wg.Done()
}

//args: callback function funcName(with or without arguments){}
func runner(funcName func()) {
	funcName()
	fmt.Println("runner finished")
}

//WaitGroup to make runner()  running background
//var wg sync.WaitGroup

func main() {
	//fmt.Println("call to haha()")
	//wg.Add(2)
	//go runner(haha)
	//fmt.Println("call to haha()")
	//
	//fmt.Println("call to tony()")
	//go runner(tony)
	//fmt.Println("call to tony()")
	//wg.Wait()
	runner(haha)
	runner(tony)


}
