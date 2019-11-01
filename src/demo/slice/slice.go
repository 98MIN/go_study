package main

import "fmt"

func main()  {
	// 切片
	_slice := []int{1,2,3,4,5}
	// 数组
	_arr := [...]int{1,2,3,4,5}
	// 切片可以扩展
	_slice = append(_slice, 6, 7, 8, 9)
	// 数组不可扩展，固定长度
	//first argument to append must be slice; have [5]int
	// _arr = append(_arr, 6 ,7 ,8 ,9)

	fmt.Println(_slice, _arr)
}
