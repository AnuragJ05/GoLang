package main

import "fmt"

func main() {

   //Type 1
   a:=40
   
   switch a{
   case 10: fmt.Println("a=10")
   case 20,30,40: fmt.Println("a is either 20,30,40")
   case 50: fmt.Println("a=50")
   default: fmt.Println("No option")
   }
   
   //Type 2
   
   b :=10
   c :=100
   
   switch{
   case (c%b != 0): fmt.Println("No")
   case (c%b == 0): fmt.Println("Yes")
   default: fmt.Println("Error")
   }
   

}


