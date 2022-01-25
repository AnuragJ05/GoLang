package main

import (
    "fmt"
	"strings"
)

func main(){

var inp string

//fmt.Println("Enter your input")
//fmt.Scanf("%s",&inp)
//fmt.Println(inp)

inp = "My name is anurag jain"

inp2 := strings.Split(inp," ")
fmt.Println(inp2)

}