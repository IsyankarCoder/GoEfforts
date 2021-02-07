package main

import "fmt"

func main(){
	messages:=make(chan string,4)

	go func(){
		messages <-"ping IsyankarCoder"
		messages <-"ddsd"
	}() 
	
	msg:= <-messages
	fmt.Println(msg)n
	fmt.Println(<-messages)	 
	fmt.Println(<-messages)
	

 
}