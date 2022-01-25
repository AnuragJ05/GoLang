package main

import "fmt"

func main(){
a :=10
b :=20

fmt.Println("a = ",a)
fmt.Println("b = ",b)

a,b = swap(a,b)
fmt.Println("a = ",a)
fmt.Println("b = ",b)

}

func swap(a int, b int) (int,int){

return b,a

}