package main

import "fmt"

func oddoreven() {
	if 7%2 == 0 {
		fmt.Println("Tek")
	} else {
		fmt.Println("Çift")
	}
	//
}

func hasdigit() {
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}

func main() {
	var x float64 = 10.8
	var y float64 = 9.87

	z := x + y

	fmt.Println(z)
	fmt.Println(x)
	fmt.Println(y)

	d := &x

	fmt.Println(d)
	fmt.Println(*d)
	fmt.Println("volki tolki")

	var k, l int = 1, 2
	fmt.Println(k, l)

	o := 1
	for o <= 3 {
		fmt.Println(o)
		o++
	}

	for u := 1; u <= 3; u++ {
		fmt.Println(u)
	}

	oddoreven()
	hasdigit()

}
