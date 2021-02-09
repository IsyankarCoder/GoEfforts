package main
import "fmt"

func printAll(v []interface{}){
  for _,val:=range v{
	  fmt.Println(val)
  }
}

func main(){
 names:=[]string{"volkan","tolkan","55"}
 vals:=make([]interface{},len(names))
 for i,v:= range names{
   vals[i]=v
 }

 printAll(vals)
}