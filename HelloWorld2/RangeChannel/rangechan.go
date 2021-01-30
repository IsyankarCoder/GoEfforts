package main

import "fmt"

func main(){
	queue:=make(chan string,40)
	queue<-"bir"
	queue<-"iki"
	queue<-"iss"
	queue<-"iddss"
	close(queue)
	
	for elem:= range queue{
		fmt.Println(elem)
	}


}