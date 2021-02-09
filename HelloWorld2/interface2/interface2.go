package main

import "fmt"

type animal interface{
  Speak() string
}

type dog struct{

}

func (d dog) Speak() string{
	return "Hwaa"
}

type cat struct{

}

func (c cat) Speak() string{
	return "Miyaww"
}

type lLama struct{

}

func(l lLama) Speak() string{
	return "????"
}

type javaDeveloper struct{

}

func (j javaDeveloper) Speak() string{
	return "Design Patterns"
}


func main(){
   animals := []animal{dog{},cat{},lLama{},javaDeveloper{}}
   for _,animal :=range animals{
	   fmt.Println(animal.Speak())
   }
}