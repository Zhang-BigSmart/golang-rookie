package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {
	demo4()
}

func demo1() {

	var m map[string]int

	m = map[string]int{"one": 1, "two": 2}

	fmt.Printf("m point: %p, value: %d\n", &m, m["one"])

	mapObj := make(map[string]int, 10)
	mapObj["1"] = 2

	fmt.Printf("mapObj len: %d\n", len(mapObj))
}

func demo2() {

	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	for k, v := range scene {
		fmt.Println(k, v)
	}

	for _, v := range scene {
		fmt.Println(v)
	}

	for k := range scene {
		fmt.Println(k)
	}
}

func demo3() {

	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	// 切片
	var slic []string

	// 将map数据遍历复制到切片中
	for k := range scene {
		slic = append(slic, k)
	}

	// 排序
	sort.Strings(slic)

	fmt.Println(slic)
}

func demo4() {
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brazil"] = 4
	scene["china"] = 960

	delete(scene, "route")

	for k, v := range scene {
		fmt.Println(k, v)
	}
}

func demo5() {

	var scene sync.Map

	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("london"))

	scene.Delete("london")

	scene.Range(func(k, v interface{}) bool {
		fmt.Print("iterate:", k, v)
		return true
	})
}
