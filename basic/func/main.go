package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	demo9()
}

func demo1() {
	fmt.Println(hypot(3, 4))
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func demo2() {
	fmt.Println(maker(2))
}

func maker(x int) (int, int) {
	return x * 1, x * 2
}

func demo3() {
	var f func()

	f = fire

	f()
}

func fire() {
	fmt.Println("fire")
}

func demo4() {
	f := func(data string) {
		fmt.Println("hello", data)
	}
	f("world")
}

func demo5() {

	// 创建一个累加器，初始值为1
	accumulator := Accmulate(1)

	fmt.Println(accumulator())

	fmt.Println(accumulator())

	fmt.Printf("%p\n", &accumulator)

	accumulator2 := Accmulate(10)

	fmt.Println(accumulator2())

	fmt.Printf("%p\n", &accumulator2)

}

func Accmulate(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func demo6() {

	a := func(args ...int) {
		for _, arg := range args {
			fmt.Println(arg)
		}

	}

	a(1, 2, 3)
}

func demo7() {
	MyPrintf(1, 22311, "hello world", 0.115)
}

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Printf("%v is an unknown type, type is %T\n", arg, arg)
		}
	}
}

func demo8() {

	fmt.Println("defer begin")

	defer fmt.Println(1)

	defer fmt.Println(2)

	defer fmt.Println(3)

	fmt.Println("defer end")
}

func demo9() {
	start := time.Now()
	for i := 0; i < 10000000; i++ {
	}
	elasped := time.Since(start)
	fmt.Println("time cost:", elasped)
}
