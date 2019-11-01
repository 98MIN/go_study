package main
import "fmt"

func pass_by_val(val int) int {
	val++

	return val
}

func main()  {
	val := 1

	fmt.Println(pass_by_val(val))

	fmt.Println(val)
}
