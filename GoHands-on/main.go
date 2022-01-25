package main

import (
	"fmt"
	"math"
	"unsafe"
	"GoLearn/calculator"
	"GoLearn/structure_pack"
	"GoLearn/structure_pack_new"
	"time"
	"sync"
	"os"
	"GoLearn/errors"
	//"runtime/debug"
	"reflect"
)

//8.Functions
func check_even_odd(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}


//9.Multiple return from Functions
func swap(num1 int, num2 int) (int, int) {
	return num2, num1
}


//17.Variadic
func variadic(arr ...int){
	fmt.Println("variadic arr = ",arr)
}


//20.Pointers
func point(val *int){
	*val =1
}


//21.structure
//22.Pointers to a struct
type structure struct{
	struct_name string
	struct_age int
}


//24.Nested structure
type structure_mark struct{
	struct_mark int
}
type structure_sub struct{
	struct_sub string
}
type structure_nested struct{
	struct_name string
	//Nested structs
	sub structure_sub
	//Promoted fields
	structure_mark
}


//25. Method
type Rectangle struct{
	length int
	width int
}
func (rt *Rectangle) set_rect_val(length int, width int){
	rt.length= length
	rt.width=width
}
func (rt Rectangle) area(){
	area:= rt.length * rt.width
	fmt.Println("Area of Rectangle is : ",area)
}


//26.Interface
type interface_example interface{
	interface_func()
}
type interface_stuct struct{
	struct_name string
	struct_age int
}
func (is interface_stuct) interface_func(){
	fmt.Println("Struct name from interface = ",is.struct_name)
	fmt.Println("Struct age from interface = ",is.struct_age)
}


//27.Interface with pointer
type interface_example2 interface{
	interface_func2()
}
type interface_stuct2 struct{
	struct_name string
	struct_age int
}
func (is *interface_stuct2) interface_func2(){
	fmt.Println("Struct name from interface = ",is.struct_name)
	fmt.Println("Struct age from interface = ",is.struct_age)
}

//28.Goroutine
func hello(){
	fmt.Println("Hello goroutine!!")
}


//29.Channel
func hello2(done chan bool){
	fmt.Println("Hello goroutine with channel!!")
	done <-true
}


//30.Goroutine and channel example
//sum of a^2 and a^3
func a_square(c chan int, a int){
	c <- a*a
}
func a_cube(c chan int, a int){
	c <- a*a*a
}


//31.Closing channels and for range loops on channels
func chan_loop(c chan int){
	for i:=1;i<=3;i++{
		c <-i
	}
	close(c)
}


//32.Buffered Channel
func write_chan(ch chan int) {  
    for i := 0; i < 3; i++ {
        ch <- i
        fmt.Println("successfully wrote", i, "to ch")
    }
    close(ch)
}


//33.WaitGroup
func process(i int, wg *sync.WaitGroup) {  
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)
    wg.Done()
}


//34.Mutex
var mutex_x  = 0  
func increment(wg *sync.WaitGroup, m *sync.Mutex) {  
    m.Lock()
    mutex_x = mutex_x + 1
    m.Unlock()
    wg.Done()   
}

//35.select
func select_chan(c chan string){
	time.Sleep(2625 * time.Millisecond)
	c <- "signal sent"
}


//37.Composition
type comp_struct1 struct{
	name string
	age int
}
func(cs1 comp_struct1) comp_method1() string{
	fmt.Printf("%s %d\n", cs1.name, cs1.age)
	return fmt.Sprintf("%s %d", cs1.name, cs1.age)
}
type comp_struct2 struct{
	bio string
	comp_struct1
}
func(cs2 comp_struct2) comp_method2(){
	fmt.Println("Details are : ",cs2.comp_method1())
	fmt.Println("Bio = ",cs2.bio)
}


//38.Polymorphism
type poly_interface interface{
	poly_method()
}
type poly_struct1 struct{
	poly1_name	string
}
func(ps1 poly_struct1) poly_method(){
	fmt.Println("From poly1 : ",ps1.poly1_name)
}
type poly_struct2 struct{
	poly2_name	string
}
func(ps2 poly_struct2) poly_method(){
	fmt.Println("From poly2 : ",ps2.poly2_name)
}
func polycalling(pi []poly_interface){
	for _,v:= range pi{
		v.poly_method()
	}
}


//39.Defer
func defer_fun(){
	defer fmt.Println("Defer fun call after defer_fun complete :stack1")
	defer fmt.Println("Defer fun call after defer_fun complete :stack2")
	fmt.Println("In defer_fun........")
}

//44.User defined function types
type add func(num1 int, num2 int) int  


//45.Returning functions from other functions
func re_fun() func (a int) int{
	f := func (a int)int {
		return a*a
	}
	return f
}

//46.Passing functions as arguments to other functions
func pas_fun(a func(num1 int, num2 int) int) {  
    fmt.Println(a(60, 7))
}

//47.Reflection
type reflection_struct struct{
	name string
	age int
}
func createquerry(q interface{}){
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	n :=v.NumField()
	fmt.Println("Reflection - Type ", t)
	fmt.Println("Reflection - Kind ", k)
	fmt.Println("Reflection - Value ", v)
	fmt.Println("Reflecting - Number of fields", n)
	fmt.Printf("Reflecting - Field:%d type:%T value:%v\n", 0, v.Field(0), v.Field(0))
	fmt.Printf("Reflecting - Field:%d type:%T value:%v\n", 1, v.Field(1), v.Field(1))

}



//49.Recover panic
func recover_panic(){
	if r:=recover();r!=nil{
		fmt.Println("recovered from ", r)
		//Getting Stack Trace after Recover
		//debug.PrintStack()
	}
}


//Main Function
func main() {

	//1.Variables
	var age = 23
	name := "anurag"
	fmt.Println("Name = ", name, " Age = ", age)
	var default_int_value int
	fmt.Println("Default value of int is ", default_int_value)
	

	//2.Initializing multiple variable
	var x, y = 2, 5
	fmt.Printf("X= %d , Y= %d\n", x, y)
	var (
		empno   = 11
		empname = "anurag"
	)
	fmt.Println("EmpNo = ", empno, " EmpName = ", empname)
	

	//3.Find minimum variable
	num1 := 100.0
	num2 := 200.0
	fmt.Println("Min no. is : ", math.Min(num1, num2))
	

	//4.Size of variables
	num3 := 20
	fmt.Printf("Type of num3 is %T , size of num3 is %d \n", num3, unsafe.Sizeof(num3))
	

	//5.Complex variables
	complex1 := complex(10, 20)
	fmt.Println("Complex no. is : ", complex1)
	

	//6.Type conversion
	//fmt.Println("Sum of num1 + num3 is ",num1+num3)
	fmt.Println("Sum of num1 + num3 is ", int(num1)+num3)
	

	//7.Constant Variables
	const num4 = 50
	fmt.Println("num4 = ", num4)
	//num4 = 500
	//const num5 = Math.Sqrt(5)
	//var num4 =500
	

	//8.Functions
	var num5 = 200
	fmt.Println("num5 = ", num5)
	if check_even_odd(num5) {
		fmt.Println("num5 is even")
	} else {
		fmt.Println("num5 is odd")
	}


	//9.Multiple return from Functions
	num6, num7 := 12, 24
	fmt.Printf("Before swap - num6= %d , num7= %d\n", num6, num7)
	num6, num7 = swap(num6, num7)
	fmt.Printf("After swap - num6= %d , num7= %d\n", num6, num7)
	

	//10.Custom Packages
	cal := calculator.Calculate(4, 2, "+")
	fmt.Println("calculate= ", cal)
	

	//11.if else statements
	if num8 := 25.0; math.Sqrt(num8) == 5{
		fmt.Println("SQRT 0f 25 is 5")
	} else{
		fmt.Println("SQRT 0f 25 is not 5")
	}


	//12.For loop
	for i := 1;i<=10;i++{
		if i==2{
			continue
		} else if i==8{
			break
		} else{
			fmt.Println("i= ",i)
		}
	}
	j:=1
	for j<=5{
		fmt.Printf("%d * %d = %d\n",10,j,10*j)
		j+=1
	}
	for k1:=0;k1<5;k1++{
		for k2:=0;k2<k1;k2++{
			fmt.Printf("*")
		}
		fmt.Println()
	}


	//13.Switch case 
	num9:=3
	switch num9{
	case 1:
		fmt.Println("num9 = 1")
	case 2,3,4:
		fmt.Println("2<=num9<=4")
	default:
		fmt.Println("num9 = 5")
	}


	//14.switch without arguments
	num10:=100
	switch{
	case num10 >=1 && num10 <=50 :
		fmt.Println("1<=num10<=50")
	case num10 >51 && num10 <=100 :
		fmt.Println("51<num10<=100")
	default:
		fmt.Println("num10>100")
	}


	//15.Array and slice
	var arr1 = [5]int{1,2,3,4}
	fmt.Println("arr1 is : ",arr1)
	var arr2 [3]int
	arr2[0],arr2[1],arr2[2]=10,20,30
	fmt.Println("arr2 is : ",arr2)
	for i,x := range arr2{
		fmt.Printf("arr2[%d] = %d\n",i,x)
	}
	arr3 := [5]int{}
	for i:=0;i<5;i++{
		arr3[i] = i*5
	}
	for i:=0;i<5;i++{
		fmt.Printf("arr3[%d] = %d\n",i,arr3[i])
	}
	fmt.Printf("arr3[1:3] = %d\n",arr3[1:3])
	arr5:=arr3[1:3]
	fmt.Printf("length of slice %d capacity %d of arr5 from arr3[1:3]\n", len(arr5), cap(arr5))
	

	//16.Append in unsized array
	var arr4 []int
	arr4 = append(arr4,98)
	arr4 = append(arr4,99)
	arr4 = append(arr4,100)
	fmt.Println("arr4 = ",arr4)
	arr4 = append(arr4,arr5...)
	fmt.Println("aar5 appended to arr4 then arr4 = ",arr4)
	

	//17.Variadic
	variadic(1,2,3)
	variadic(1,2,3,4,5)
	variadic(1,2,3,4,5,6,7,8)


	//18.Map
	mp := make(map[string]int)
	mp["Anurag"] =100
	mp["Prince"] =98
	fmt.Println("MP = ",mp)
	mp2 :=map[string]string{
		"Anurag":"Jain",
	}
	mp2["Prince"]="jain"
	for k,v:=range mp2{
		fmt.Printf("mp2[%s] : %s\n",k,v)
	}
	fmt.Printf("mp[Calsoft] : %d\n",mp["Calsoft"])
	delete(mp,"Anurag")
	fmt.Println("MP after deleting Anurag= ",mp)
	

	//19.string and runes
	char1:="señor"
	fmt.Printf("ñ of señor in char: %c\n",char1[2])
	runes := []rune(char1)
	fmt.Printf("ñ of señor in char by using rune: %c\n",runes[2])
	str1 := "Anurag"
	str2 := "Jain"
	str3 := str1+" "+str2
	fmt.Println("STR1+STR2= ",str3)
	// str3[0]='P'
	// fmt.Println("STR3 when first letter is change= ",str3)
	runes2 := []rune(str3)
	runes2[0]='P'
	fmt.Println("STR3 when first letter is change after using runes= ",string(runes2))


	//20.Pointers
	var ptr1 *int
	val1 := 100
	fmt.Println("ptr1 before initialization is : ",ptr1)
	ptr1 = &val1
	fmt.Println("ptr1 after initialization is : ",ptr1)
	fmt.Println("Value of ptr1 after initialization is : ",*ptr1)
	point(&val1)
	fmt.Println("Value of val1 after pointer pass a address in function is : ",val1)


	//21.structure
	struct1 := structure{"Anurag",23}
	fmt.Println("Struct1 = ",struct1)
	var struct2 structure
	struct2.struct_name ="Prince"
	fmt.Println("struct2.struct_name = ",struct2.struct_name)
	fmt.Println("struct2.struct_age = ",struct2.struct_age)
	

	//22.Pointers to a struct
	struct3 := &structure{
		struct_name:"Anurag",
		struct_age:23,
	}
	fmt.Println("Struct3_struct_name = ",(*struct3).struct_name)


	//23.Exported structs and fields
	struct4 :=structure_pack.Structure_pack_details{
		Name: "Anurag jain",
		Age: 23,
	}
	fmt.Println("struct4 = ",struct4)


	//24.Nested structure
	struct5 := structure_nested{
		struct_name: "Anurag",
		sub:structure_sub{
			struct_sub: "GoLang",
		},
		structure_mark:structure_mark{
			struct_mark: 99,
		},
	}
	fmt.Println("struct5 with nested struct : ",struct5)


	//25. Method
	var struct6 Rectangle
	struct6.set_rect_val(10,15)
	struct6.area()


	//26.Interface
	inter_is := interface_stuct{"Anurag",23}
	var inter interface_example
	fmt.Printf("Interface type %T value %v before structure type initialization\n", inter, inter)
	inter=inter_is
	fmt.Printf("Interface type %T value %v after structure type initialization\n", inter, inter)
	inter.interface_func()


	//27.Interface with pointer
	inter_is2 := interface_stuct2{"Prince",23}
	var inter2 interface_example2
	//fmt.Printf("Interface type %T value %v before structure type initialization\n", inter2, inter2)
	inter2=&inter_is2
	//fmt.Printf("Interface type %T value %v after structure type initialization\n", inter2, inter2)
	inter2.interface_func2()


	//28.Goroutine
	go hello()
	//If the main Goroutine terminates then the program will be terminated and no other Goroutine will run.
	//so, timer is added
	time.Sleep(1 * time.Second)
	fmt.Println("Hello main!!")

	//29.channel
	done := make(chan bool)
	go hello2(done)
	<-done
	fmt.Println("Hello main with channel!!")

	//30.Goroutine and channel example
	//sum of a^2 and a^3
	asquare :=  make(chan int)
	acube :=  make(chan int)
	a := 2
	go a_square(asquare,a)
	go a_cube(acube,a)
	val1,val2 := <-asquare, <-acube
	fmt.Println("sum of square and cube of a=2 using goroutine and channel is :",(val1+val2))

	//31.Closing channels and for range loops on channels
	chan1 := make(chan int)
	go chan_loop(chan1)
	for{
		v,ok := <-chan1
		if ok == false{
			break
		}
		fmt.Println("Received ", v, ok)
	}
	chan2 := make(chan int)
	go chan_loop(chan2)
	for v := range chan2{
		fmt.Println("Received ", v)
	}


	//32.Buffered Channel
	//set capacity 2
	chan3 := make(chan int,2)
	go write_chan(chan3)
	time.Sleep(2 * time.Second)
	for v := range chan3{
		fmt.Println("read value", v,"from ch")
		time.Sleep(2 * time.Second)
	}


	//33.WaitGroup
	no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
        wg.Add(1)
        go process(i, &wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")

	//34.Mutex
	var w sync.WaitGroup
    var m sync.Mutex
    for i := 0; i < 1000; i++ {
        w.Add(1)        
        go increment(&w, &m)
    }
    w.Wait()
    fmt.Println("final value of mutex_x", mutex_x)

	//35.select
	chan4 := make(chan string)
	go select_chan(chan4)
	flag :=0
	for flag==0{
		time.Sleep(1000 * time.Millisecond)
		select{
		case v := <-chan4:
			fmt.Println("received value: ", v)
			flag =1
		default:
			fmt.Println("Waiting for signal to sent")
		}
	}


	//36.New() function
	spn := structure_pack_new.New("Anurag",23)
	spn.Display()


	//37.Composition
	cs1 := comp_struct1{"Anurag",23}
	cs2 := comp_struct2{"Bio of Anurag.....",cs1}
	cs2.comp_method2()


	//38.Polymorphism
	ps1 := poly_struct1{"poly_struct1_name"}
	ps2 := poly_struct2{"poly_struct2_name"}
	pi := []poly_interface{ps1,ps2}
	polycalling(pi)


	//39.Defer
	defer fmt.Println("\n\n\n\n\n\nDefer after main func complete....")
	defer_fun()


	//40.Error Handling
	f, err := os.Open("/test.txt")
	if err != nil{
		fmt.Println("Error is : ",err)
	} else{
		fmt.Println(f.Name(), " opened successfully")
	}


	//41.Custom error
	radius := -100
	if radius<0{
		cus_err := errors.New("radius is less than 0")
		if cus_err != nil{
			fmt.Println("Custom error is : ",cus_err)
		}
	}

	//42.Anonymous functions
	anonymous_func := func(){
		fmt.Println("Anonymous functions")
	}
	anonymous_func()
	fmt.Printf("%T\n", anonymous_func)

	//43.Closure and Anonymous functions without calling
	closure_val :=100
	func() {
        fmt.Println("hello world first class function and closure_val is : ",closure_val)
    }()
	

	//44.User defined function types
	var ad add = func(num1 int, num2 int) int {
        return num1 + num2
    }
    su := ad(5, 6)
    fmt.Println("Sum", su)


	//45.Returning functions from other functions
	re_func := re_fun()
	fmt.Println("2*2= ",re_func(2))


	//46.Passing functions as arguments to other functions
	pass_func := func(num1 int, num2 int) int {
        return num1 + num2
    }
    pas_fun(pass_func)


	//47.Reflection
	ref_st := reflection_struct{"Anurag",23}
	createquerry(ref_st)


	//48.Panic
	// defer fmt.Println("Defer call in panic")
	// fName1 := ""
	// if fName1 == ""{
	// 	panic("fName1 cannot be nil")
	// }
	// fmt.Println("fName1 after panic is : ",fName1)


	//49.Recover panic
	defer recover_panic()
	fName2 := ""
	if fName2 == ""{
		panic("fName2 cannot be nil")
	}
	fmt.Println("fName2 after panic is : ",fName2)

}