package main

import "fmt"

func main(){
a :=5

fac := fact(a)
fmt.Println(fac)
}


func fact(a int)int{

if a ==0{
return 1
}else{
return a * fact(a-1)
}
}