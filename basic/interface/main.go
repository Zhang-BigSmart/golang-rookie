package main

import "fmt"

func main() {
	demo3()
}

func demo1() {
	var any interface{}

	any = 1
	fmt.Println(any)

	any = "hello"
	fmt.Println(any)

	any = false
	fmt.Println(any)
}

func demo2() {
	var a int = 1

	var i interface{} = a

	var b int = i.(int)

	fmt.Println(b)
}

func demo3() {
	factory := &CarFactory{"Car"}
	doProduce(factory)
	doConsume(factory)
}

type Factory interface {
	Produce() bool
	Consume() bool
}

type CarFactory struct {
	ProductName string
}

func (c *CarFactory) Produce() bool {
	fmt.Printf("CarFactory生产%s成功\n", c.ProductName)
	return true
}

func (c *CarFactory) Consume() bool {
	fmt.Printf("CarFactory消费%s成功\n", c.ProductName)
	return true
}

func doProduce(factory Factory) bool {
	return factory.Produce()
}

func doConsume(factory Factory) bool {
	return factory.Consume()
}
