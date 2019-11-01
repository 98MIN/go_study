package main

import "fmt"

func main(){
	score := 80

	/*
	* switch 不需要添加break
	*/
	// switch {
	// 	case score > 90:
	// 		fmt.Println("A")
	// 	case score > 75:
	// 		fmt.Println("B")
	// 	case score > 60:
	// 		fmt.Println("C")
	// 	case score < 60:
	// 		fmt.Println("D")
	// 	default:
	// 		fmt.Println("请输入合法成绩")
	// }


	/**
	* 使用了fallthrough之后不会去判断下一个case是否为true直接执行下一行语句。
	* 下一行未使用fallthrough则执行完这条case之后退出；使用了则继续下一行
	*/
	// 	switch {
	// 		case score > 90:
	// 			fmt.Println("A")
	// 			fallthrough
	// 		case score > 75:
	// 			fmt.Println("B")
	// 			fallthrough
	// 		case score > 60:
	// 			fmt.Println("C")
	// 			fallthrough
	// 		case score < 60:
	// 			fmt.Println("D")
	// 			fallthrough
	// 		default:
	// 			fmt.Println("请输入合法成绩")
	// }

	switch {
		case score > 90:
		case score > 75:
			fmt.Println("A")
		case score > 60:
			fmt.Println("C")
			fallthrough
		case score < 60:
			fmt.Println("D")
			fallthrough
		default:
			fmt.Println("请输入合法成绩")
	}
}
