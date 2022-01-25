package main

import (
    "fmt"
)

func main(){

var arr [10] int
var i int

for i=0;i<10;i++{
arr[i]=i+1
}

for i=0;i<10;i++{
fmt.Println(arr[i])
}

}