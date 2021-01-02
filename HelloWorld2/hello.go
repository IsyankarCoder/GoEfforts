package main

import "fmt"

func main(){
	var x float64=10.8;	
	var y float64=9.87;

	z:= x+y;

	fmt.Println(z);
	fmt.Println(x);
	fmt.Println(y);
	
	
	d:=&x;

	fmt.Println(d);
	fmt.Println(*d);
	fmt.Println("volki tolki");
}