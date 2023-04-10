package main

import (
	"container/list"
	"fmt"
)

func main() {
	demo2()
}

func demo1() {
	l := list.New()

	// 尾部添加
	l.PushBack("first")
	// 头部添加
	l.PushFront(67)

	element := l.PushBack("can")

	l.InsertAfter("do", element)

	l.InsertBefore("us", element)

	l.Remove(element)

	fmt.Printf("l value: %v \n", l)
	fmt.Printf("element: %v, %T", element, element)
}

func demo2() {

	l := list.New()

	l.PushBack("back")

	l.PushFront("test")

	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
