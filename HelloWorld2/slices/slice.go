package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "sıfır"
	s[1] = "bir"
	s[2] = "iki"

	fmt.Println("dolu", s)
	fmt.Println("len", len(s))

	s = append(s, "d")
	s = append(s, "e", "f0")

	fmt.Println("apd:", s)


	c:= make([]string,len(s))
	copy(c,s)
	fmt.Println("cpoy",c)

	l:=s[2:3]
	fmt.Println("ll	:",l)

	l=s[:0]
	fmt.Println("l5",l)

	l=s[3:]
	fmt.Println("l5",l)

	t:=[] string{"a","b","c"}
	fmt.Println(t)


	twoD:=make([][]int,3)
	for i:=0;i<3;i++{
		innerLen:=i+1
		twoD[i]= make([]int, innerLen)
		for j:=0;j<innerLen;j++{
				twoD[i][j]=i*j
		}
	}

	fmt.Println(twoD)

}
