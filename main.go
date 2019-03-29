package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var x int
	x = "string" // ver.1.9だとgo vetでエラー吐かない
}
