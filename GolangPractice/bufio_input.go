package main

import (
		"fmt"
		"bufio"
		"os"
		"strings"
)

func main(){

scanner := bufio.NewScanner(os.Stdin)

fmt.Println("Enter input :")
scanner.Scan()
val := scanner.Text()
fmt.Println(val)


val_split := strings.Split(val," ")
fmt.Println(val_split)
}