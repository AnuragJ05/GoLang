package main

import "fmt"

func countBits(num int) int {
    //fmt.Println(num)
    count :=0
    s := fmt.Sprintf("%b",num)
    fmt.Println(s)
    for _,j:= range s{
        if j == '1'{
            count++
        } 
    }
    return count

}

func main(){

num := 12

fmt.Println(countBits(num))

}