package main

import "fmt"

func main() {
	demo5()
}

func demo1() {
	source := [...]int{1, 2, 3}
	sli := source[0:1]

	fmt.Printf("sli value is %v\n", sli)
	fmt.Printf("sli len is %v\n", len(sli))
	fmt.Printf("sli cap is %v\n", cap(sli))
}

func demo2() {

	arr1 := [...]int{1, 2, 3, 4}
	arr2 := [...]int{1, 2, 3, 4}

	sli1 := arr1[0:2]
	sli2 := arr2[2:4]

	fmt.Printf("sli1 pointer: %p, len: %v, cap: %v, value: %v\n", &sli1, len(sli1), cap(sli1), sli1)
	fmt.Printf("sli2 pointer: %p, len: %v, cap: %v, value: %v\n", &sli2, len(sli2), cap(sli2), sli2)

	newSli1 := append(sli1, 5)

	fmt.Printf("newSli1 pointer: %p, len: %v, cap: %v, value: %v\n", &newSli1, len(newSli1), cap(newSli1), newSli1)
	fmt.Printf("source arr1 become %v\n", arr1)

	newSli2 := append(sli2, 5)

	fmt.Printf("newSli2 pointer: %p, len: %v, cap: %v, value: %v\n", &newSli2, len(newSli2), cap(newSli2), newSli2)
	fmt.Printf("source arr2 become %v\n", arr2)

}

func demo4() {

	arr1 := []int{1, 2, 3}
	// 删除第一个元素，也不能说删除
	sli1 := arr1[1:]
	fmt.Printf("sli1 pointer: %p, value: %v\n", &sli1, sli1)
	// 删除第一个元素
	sli2 := append(arr1[:0], arr1[1:]...)
	fmt.Printf("sli2 pointer: %p, value: %v\n", &sli2, sli2)

	// slic3 := arr1[:copy(arr1, arr1[1:])]

}
