package main

import (
		"fmt"
		//"strings"
		"sort"
)


func main(){

s :=[]string{"zbc","asx","jst"}
sort.Strings(s)
fmt.Println(s)



i := []int{34,42,1,3,2,452,6,8,5}

fmt.Println(sort.IntsAreSorted(i))
sort.Ints(i)
fmt.Println(i)
fmt.Println(sort.IntsAreSorted(i))

}