package main

import (
	"fmt"
	"time"
)
       
func main(){
   i:=2

   fmt.Println("write", i, "as") 
   today:= time.Now().Weekday()
   fmt.Println(today)
    
   switch time.Now().Weekday() {
   case time.Saturday,time.Sunday: 
           fmt.Println("Haftasonu tatlım");
   default:
      fmt.Println("Aq. iş var");
   }

   t:=time.Now()
   fmt.Println(t)

   switch {
   case t.Hour() < 12:
       fmt.Println("It's before noon")
   default:
       fmt.Println("It's after noon")
   }

   whatAmI:= func(i interface{}){
      switch i.(type) {
      case bool: 
         fmt.Println("Boolmuş bu")
      case int:
         fmt.Println("Integer lo")
      default:
         fmt.Println("ne idügü belirsiz aq.",i)         
      }


   }

   whatAmI(true)
   whatAmI(55)
   whatAmI("volki tolki")
}