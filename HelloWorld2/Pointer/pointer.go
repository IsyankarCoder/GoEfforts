package main

import "fmt"



func zeroVal(val int){
  val=0
}

func zeroptr(intptr *int){
  *intptr=0
}

func main(){
  i:=1
  fmt.Println("initial : ",i)

  zeroVal(i)
  fmt.Println("zeroVal : ",i)

  zeroptr(&i)
  fmt.Println("zeroptr",i)
 
  fmt.Println("pointer:",&i)
}