package main

import "fmt"

func main() {


for i:=1;i<=5;i++ {
fmt.Println(i)
}


a:=1
b:=5
for a<=b{
fmt.Println("a is ",a)
a++
}


list := [6]int{11,2,13,4}

for i,j:= range list{
fmt.Println(i,j)
} 


for _,i:= range list{
fmt.Println(i)
} 
  
}


