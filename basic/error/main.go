package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	demo1()
}

func demo1() {
	fmt.Println("main begin")

	defer func() {
		fmt.Println("defer begin")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("defer end")
	}()
	test()
	fmt.Println("main end")
}

func test() {
	fmt.Println("test begin")
	panic("error")
	fmt.Println("test end")
}

func demo2() {
	result, error := Sqrt(-13)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func Sqrt(f float64) (float64 error) {
	if f < 0 {
		return -1, errors.New("math: square root of negative number")
	}
	return math.Sqrt(f), nil
}
