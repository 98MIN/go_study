package main

import "fmt"

func pass_by_ref(arr *[5]int) *[5]int{
	arr[0] = 200

	return arr
}

func main()  {
  arr := [...]int{1,2,3,4,5}

	pass_by_ref(&arr)
	fmt.Println(arr)
}
