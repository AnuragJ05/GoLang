package main

import (
	"fmt"
	// "strings"
	// "encoding/json"
	// "sort"
)

// type test struct {
// 	name string
// 	age  int
// }

// func (test_ob test) display() {
// 	fmt.Println("Test: ", test_ob)
// 	fmt.Println("Name: ", test_ob.name)
// 	fmt.Println("Age: ", test_ob.age)
// }

func main() {

	// tst := test{"Anurag", 23}
	// fmt.Println("Name: ", tst.name)
	// fmt.Println("Age: ", tst.age)

	// var tst2 test
	// tst2.name = "Jain"
	// tst2.age = 100

	// tst2.display()

	// var s1 string
	// var s2 string
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Println("Enter string1")
	// // fmt.Scan("%s", &s1)
	// scanner.Scan()
	// s1 = scanner.Text()
	// fmt.Println("Enter string1")
	// // fmt.Scan("%s", &s2)
	// scanner.Scan()
	// s2 = scanner.Text()
	// fmt.Println(s1 + s2)

	// var num1 int
	// fmt.Println("Enter num1")
	// fmt.Scanf("%d\n", &num1)
	// var num2 int
	// fmt.Println("Enter num2")
	// fmt.Scanf("%d\n", &num2)
	// fmt.Println("num1+num2 = ", num1+num2)

	// var arr = [5]int{1, 2, 3, 4, 5}
	// for i, x := range arr {
	// 	fmt.Printf("Value at index %d is %d\n", i, x)
	// }

	// var arr2 [5]int
	// for i := 0; i < 5; i++ {
	// 	//arr2[i] = i + 1
	// 	fmt.Printf("Enter element for index %d: ", i)
	// 	fmt.Scanf("%d\n", &arr2[i])
	// }
	// for i, x := range arr2 {
	// 	fmt.Printf("Value at index %d is %d\n", i, x)
	// }
	// fmt.Println("Array is ", arr2)
	// reverse(arr2[:])
	// fmt.Println("Array after Reverse : ", arr2[:])
	// sort.Ints(arr2[:])
	// fmt.Println("Array after sorting : ", arr2)

	// mp := map[string]string{}

	// mp["First Name"] = "Calsoft"
	// mp["Last Name"] = "Jain"
	// mp["First Name"] = "Anurag"

	// fmt.Println("Map : ", mp)
	// for key, value := range mp {
	// 	fmt.Println(key, ":", value)
	// }
	// delete(mp, "Last Name")
	// fmt.Println("Map after delete operation : ", mp)

	// js, err := json.Marshal(mp)
	// if err != nil {
	// 	fmt.Println("Encoding fail")
	// } else {
	// 	fmt.Println("Encoded data : ")
	// 	fmt.Println(js)
	// 	fmt.Println("Decoded data : ")
	// 	fmt.Println(string(js))
	// }

	// var arr []string
	// arr = append(arr, "Anurag")
	// arr = append(arr, "Jain")
	// arr = append(arr, "Calsoft")
	// arr = append(arr, "Pune")
	// fmt.Println(arr)
	// fmt.Println("Array from index 1 to 3 : ", arr[1:])
	// arr = append(arr[:1], arr[2:]...)
	// fmt.Println("Array after removing element from index 1 : ", arr)
	// joinarr := strings.Join(arr, " ")
	// fmt.Println("Array after join : ", joinarr)

	// a,b := 100,200
	// if a > b {
	// 	fmt.Println("a>b")
	// } else if a < b {
	// 	fmt.Println("a<b")
	// } else {
	// 	fmt.Println("a=b")
	// }
	// a, b = b, a
	// switch {
	// case a > b:
	// 	fmt.Println("a>b")
	// case a < b:
	// 	fmt.Println("a<b")
	// default:
	// 	fmt.Println("a=b")
	// }

	// str := " Calsoft Pvt Ltd"
	// fmt.Println("String before split : ", str)
	// arr_str := strings.Split(str, " ")
	// fmt.Println("String after split : ", arr_str)

	c1 := make(chan int)
	c2 := make(chan int)
	go go_test(c1, 5)
	go go_test(c2, 7)

	for {
		select {
		case msg, open := <-c1:
			if !open {
				break
			}
			fmt.Println(msg)
		case msg, open := <-c2:
			if !open {
				break
			}
			fmt.Println(msg)
		}

	}
}

// func reverse(arr []int) {
// 	n := len(arr)
// 	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
// 		arr[i], arr[j] = arr[j], arr[i]
// 	}

// }

func go_test(c chan int, value int) {
	for i := 1; i <= 10; i++ {
		// fmt.Println(value, " * ", i, " = ")
		c <- value * i
	}
	close(c)
}
