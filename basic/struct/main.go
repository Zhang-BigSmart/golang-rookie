package main

import (
	"fmt"
)

func main() {
	demo5()
}

func demo1() {

	//var student Student
	student := new(Student)

	name := "Edison"
	age := 18
	version := 1

	student.Name = name
	student.age = age
	student.version = &version

	name = "Adm"
	age = 28
	version = 2

	fmt.Printf("name: %s, age: %d, version %d\n", student.Name, student.age, *student.version)
}

type Student struct {
	Name    string
	age     int
	version *int
}

func demo2() {

	relation := &People{
		name: "grandfather",
		child: &People{
			name: "father",
			child: &People{
				name: "son",
			},
		},
	}

	fmt.Println(relation)
}

type People struct {
	name  string
	child *People
}

func demo3() {
	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}
	fmt.Println(addr)
}

type Address struct {
	Province    string
	City        string
	ZipCode     int
	PhoneNumber string
}

func demo4() {
	msg := &struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}
	printMsgType(msg)
}

func printMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n", msg)
}

type Factory interface {
	Produce() bool
	Consume() bool
}

type CarFactory struct {
	ProductName string
}

func (c *CarFactory) Produce() bool {
	fmt.Printf("CarFactory生产%s成功", c.ProductName)
	return true
}

func (c *CarFactory) Consume() bool {
	fmt.Printf("CarFactory消费%s成功", c.ProductName)
	return true
}

func demo5() {
	factory := &CarFactory{"Car"}
	doProduce(factory)
	doConsume(factory)
}

func doProduce(factory Factory) bool {
	return factory.Produce()
}

func doConsume(factory Factory) bool {
	return factory.Consume()
}
