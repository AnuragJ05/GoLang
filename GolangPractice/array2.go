package main

import (
    "fmt"
)

func main(){

var arr =[5] int{1,3,5,7,9}
var i int

for i=0;i<5;i++{
fmt.Println(arr[i])
}


for i,j:= range arr{
fmt.Println(i,j)
}

}