package main

import "fmt"

func main(){
	fmt.Println("hello world")
	x,y := 1,3
	var p = [2]int{x,y}
	arr := p
	fmt.Println(&arr)
	fmt.Println(p)

	a:=1
	if a>1{
		fmt.Println("hhh")
	}else{
		fmt.Println("aaa")
	}
}
