package main

import (
	"fmt"
	"time"
)

func main() {
	demo1()
}

func demo1() {

	go running()

	var input string
	fmt.Scan(&input)

}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)

		time.Sleep(time.Second)
	}

}
