package main

import "fmt"

func main(){
a :=10
b :=20

sum2 := func(a int,b int)int{
return a+b
}

fmt.Println(sum2(a,b))

}