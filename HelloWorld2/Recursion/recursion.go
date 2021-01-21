package main

import "fmt"


func Fact(n int)int{
	if n==0 {return 1}
	return n* Fact(n-1)
}

func main(){
   fmt.Println(Fact(5))
}