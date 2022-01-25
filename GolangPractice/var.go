
package main
import "fmt"

func main(){

/*
var a,b int
.\var.go:7:7: b declared but not used
*/

var a int
var c float32

a=20.0
fmt.Println("Value of a is ",a)
fmt.Printf("Type of a is %T\n",a)

c=20.0
fmt.Println("Value of c is ",c)
fmt.Printf("Type of c is %T\n",c)

b := "Anurag"
fmt.Println("Value of b is ",b)
fmt.Printf("Type of b is %T\n",b)

var d,e,f = 12,'a',"jain"
fmt.Println("Value of d is ",d)
fmt.Printf("Type of d is %T\n",d)
fmt.Println("Value of e is ",e)
fmt.Printf("Type of e is %T\n",e)
fmt.Println("Value of f is ",f)
fmt.Printf("Type of f is %T\n",f)

a = 4.0
fmt.Println("Value of a is ",a)
fmt.Printf("Type of a is %T\n",a)


const g int = 400
fmt.Println("Value of g is ",g)
fmt.Printf("Type of g is %T\n",g)

//g=500


}