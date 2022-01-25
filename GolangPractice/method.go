package main

import "fmt"


type student struct{

name string
roll_no int
marks float64
}

func(st student) display(){
fmt.Println(st)
fmt.Println(st.name)
fmt.Println(st.roll_no)
fmt.Println(st.marks)


}

func main(){
st := student{"anurag",32,90.2}
st.display()
}