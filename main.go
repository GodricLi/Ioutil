package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	//使用ioutil可以全部一次性读取文件
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed,err:", err)
	}
	fmt.Println(string(content))
}