package main

import "fmt"

func initialize_seq(n int) func() int{
i := n

return func() int{
i+=1
return i
}
}


func main(){
n :=10
next_seq:=initialize_seq(n)
fmt.Println(next_seq())
fmt.Println(next_seq())
fmt.Println(next_seq())

n=0
next_seq2:=initialize_seq(n)
fmt.Println(next_seq2())
fmt.Println(next_seq2())
fmt.Println(next_seq2())


}