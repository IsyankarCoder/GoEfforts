package main

import (
	"fmt"
	"math"
)

type geometry interface{
	area() float64
	perim() float64
} 

type rec struct{
	width,heigth float64
}

type circle struct{
	radius float64
}

func(r rec) area() float64{
	return r.width *r.heigth
}

func(r rec) perim() float64{
	return 2*r.width+2*r.heigth
}

func(c circle) area() float64{
	return  math.Pi*c.radius*c.radius
}

func(c circle) perim() float64{
	return 2*math.Pi*c.radius
}

func measure(geo geometry) {
	fmt.Println(geo)
	fmt.Println(geo.area())
	fmt.Println(geo.perim())
}

func main(){
	r:=rec{width:10,heigth:23}
	c:=circle{radius:22}
	measure(r)
	measure(c)
}
 