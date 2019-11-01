package main
import (
	"fmt"
	"io/ioutil"
)

func main(){
	/*
	* 读取文件
	*/
	content, err := ioutil.ReadFile("hello.txt")
	if(err != nil){
		fmt.Println(err)
	}
	fmt.Println(string(content))
}
