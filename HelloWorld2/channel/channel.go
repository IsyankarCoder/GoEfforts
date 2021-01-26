package main

import "fmt"

func main(){
	messages:=make(chan string,4)

	go func(){
		messages <-"ping volkiş"
		messages <-"ddsd"
	}()

 
	fmt.Println(<-messages)	 
	fmt.Println(<-messages)
}