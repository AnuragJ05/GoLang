package main

import "fmt"

func main(){
a :=10
b :=20

fmt.Println("a = ",a)
fmt.Println("b = ",b)

swap(&a,&b)

fmt.Println("a = ",a)
fmt.Println("b = ",b)

}

func swap(a *int, b *int){

temp := *a
*a = *b
*b = temp

}