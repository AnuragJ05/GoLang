//36.New() function
package structure_pack_new

import "fmt"

//donot export as first letter of structure is small
type structure_pack_new_details struct{
	name string
	age int
}

func New(name string, age int) structure_pack_new_details{
	e := structure_pack_new_details{name,age}
	return e
}

func (e structure_pack_new_details)Display(){
	fmt.Printf("name = %s, age = %d \n", e.name, e.age)
}