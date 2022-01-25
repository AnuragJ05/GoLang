package main

import "fmt"


func main(){

li := []int{1,2,3,4}
fmt.Println(li)

li=append(li,5,6)

for i,j:= range li{
fmt.Println(i,j)
}


mp := map[string]string{"abc":"vwx","bcd":"wxy","cde":"xyz"}
fmt.Println(mp)
for i,j:= range mp{
fmt.Println(i,j)
}

mp["anurag"]="jain"

delete(mp,"bcd")

fmt.Println(mp)
for i:= range mp{
fmt.Println(i,mp[i])
}
}