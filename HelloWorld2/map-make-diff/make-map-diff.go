package main

import (
	"fmt"
	"testing"
)

func benchmarktestingwithmake(b *testing.B){
  m0:= make(map[int]int,b.N)
   for i:=0;i<b.N;i++{
	   m0[i]=10000
   }
}

func bencmarkwithliteral(b *testing.B){
	m0:=map[int]int{}
	for i:=0;i<b.N;i++{
		m0[i]=10000
	}
}

func main(){
  bwm := testing.Benchmark(benchmarktestingwithmake)
  fmt.Println(bwm)

  blm:= testing.Benchmark(bencmarkwithliteral)
  fmt.Println(blm)

}
