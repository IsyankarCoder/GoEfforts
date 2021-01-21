package main

import "fmt"


func sum(nums ...int){
	fmt.Print(nums," ")
	total:=0
	for _,num := range nums{
		total+=num
	}

	fmt.Println(total)
}

func main(){
 sum(1,23)
 sum(1,23,3)

 nums:= []int{3,8,6,5}
 sum(nums...)
}