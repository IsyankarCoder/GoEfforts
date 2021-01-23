package main

import "fmt"

type person struct{
	name string
	age  int
}

func newPerson(name string) *person{
	p:=person{name:name}
	p.age=42
	return  &p
}

func main(){
   fmt.Println(person{"Bob",20})
   fmt.Println(person{name:"Volkan",age:55})
   fmt.Println(person{name: "Fred"})
   fmt.Println(&person{name: "Ann", age: 40})
   fmt.Println(newPerson("Jon").name)


   s:=person{name:"volkan",age:42}
   fmt.Println(s.name)

   sp:=&s
   fmt.Println(sp.age)

   sp.age=38
   fmt.Println(sp.age)
}