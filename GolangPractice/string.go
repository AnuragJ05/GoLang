package main

import (
    "fmt"
    "strings"
)

func main(){

var s = "Anurag"

fmt.Println("string is ",s)
fmt.Println("len of string is ",len(s))

s2 := "Anurag jain"

fmt.Println("string is ",s2)
fmt.Println("len of string is ",len(s2))


fmt.Println("first char by println is ",s2[0])
fmt.Printf("first char by printf is %c\n",s2[0])
fmt.Println("first 2 char is ",s2[0:2])

fmt.Println("s+s2 is - ",s+s2)

li :=  []string{"Hello","world!"}
fmt.Println("join list is - ",strings.Join(li," "))
}