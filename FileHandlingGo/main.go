package main

import(
	"fmt"
	"io/ioutil"
	"flag"
	"github.com/gobuffalo/packr/v2"
	"bufio"
	"os"
	"sync"
)

func main(){
	//Opening the file that not exists
	fmt.Println("Opening the file that not exists")
	data1,err1 := ioutil.ReadFile("text2.txt")
	if err1!= nil{
		fmt.Println("Error in opening file")
	} else{
		fmt.Println("Data from file : ",string(data1))
	}


	//Opening the existing file with aboslute/relative path
	fmt.Println("\nOpening the existing file with aboslute/relative path")
	data2,err2 := ioutil.ReadFile("text1.txt")
	if err2!= nil{
		fmt.Println("Error in opening file")
	} else{
		fmt.Println("Data from file : ",string(data2))
	}


	//Passing the file path as a command line flag
	//go run main.go -fpath=..\FileHandlingGo\text1.txt
	fmt.Println("\nPassing the file path as a command line flag")
	fptr := flag.String("fpath", "text1.txt", "file path to read from")
    flag.Parse()
    fmt.Println("value of fpath is", *fptr)
	data3,err3 := ioutil.ReadFile(*fptr)
	if err3!= nil{
		fmt.Println("Error in opening file")
	} else{
		fmt.Println("Data from file : ",string(data3))
	}


	//Bundling the text file along with the binary
	//go get -u github.com/gobuffalo/packr/v2/...  
	fmt.Println("\nBundling the text file along with the binary")
	box := packr.New("fileBox","../FileHandlingGo")
	data4,err4 := box.FindString("text1.txt")
	if err4!= nil{
		fmt.Println("Error in opening file")
	} else{
		fmt.Println("Data from file : ",string(data4))
	}



	//Reading a file in small chunks
	fmt.Println("\nReading a file in small chunks of 30 byte")
	f5,err5 := os.Open("text1.txt")
	if err5 !=nil{
		fmt.Println("Error in opening file")
	} else{
		defer f5.Close()
		r5 := bufio.NewReader(f5)
		b5 := make([]byte,30)
		for{
			//fmt.Println("b5= ",b5)
			//fmt.Println("r5= ",r5)
			data5,err := r5.Read(b5)
			if err != nil {
				fmt.Println("Error reading file:", err)
				break
			}
			fmt.Println(string(b5[0:data5]))
		}
	}

	//Reading a file line by line
	fmt.Println("\nReading a file line by line")
	f6,err6 := os.Open("text1.txt")
	if err6 !=nil{
		fmt.Println("Error in opening file")
	} else{
		defer f6.Close()
		r6 := bufio.NewScanner(f6)
		for r6.Scan(){
			fmt.Println(r6.Text())
		}
	}


	//Creating and writing in a file
	fmt.Println("\nCreating a file")
	f7,err7 := os.Create("text3.txt")
	if err7 != nil{
		fmt.Println("Error in creating file")
	}else{
		defer f7.Close()
		fmt.Println("file created successfully!!!")
		fmt.Println("\nWriting to a file.......")
		len7,err := f7.WriteString("Hello World")
		if err !=nil{
			fmt.Println("Error in writing in a file ",err)
		} else{
			fmt.Println(len7," byte written successfully...")	
		}
	}


	//writing bytes in a file
	fmt.Println("\nCreating a file")
	f8,err8 := os.Create("text4.txt")
	if err8 != nil{
		fmt.Println("Error in creating file")
	}else{
		defer f8.Close()
		fmt.Println("file created successfully!!!")
		fmt.Println("\nWriting bytes to a file.......")
		d8 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
		len8,err := f8.Write(d8)
		if err !=nil{
			fmt.Println("Error in writing in a file ",err)
		} else{
			fmt.Println(len8," byte written successfully...")	
		}
	}


	//writing line by line in a file
	fmt.Println("\nCreating a file")
	f9,err9 := os.Create("text5.txt")
	if err9 != nil{
		fmt.Println("Error in creating file")
	}else{
		fmt.Println("file created successfully!!!")
		fmt.Println("\nWriting line by line to a file.......")
		d9 := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}
		for _,v := range d9{
			len9,err := fmt.Fprintln(f9,v)
			if err !=nil{
				fmt.Println("Error in writing in a file ",err)
			}else{
				fmt.Println(len9," byte written successfully...")	
			}
		}
	}
	//close the file after use
	err9c := f9.Close()
    if err9c != nil {
        fmt.Println(err9c)
    }


	//append line in a file
	fmt.Println("\nappend line in a file.......")
	f10,err10 := os.OpenFile("text5.txt",os.O_APPEND|os.O_WRONLY, 0644)
	if err10 != nil{
		fmt.Println("Error in opening file")
	}else{
		defer f10.Close()
		d10 := "File handling is easy."
		len10,err := fmt.Fprintln(f10,d10)
		if err !=nil{
			fmt.Println("Error in writing in a file ",err)
		}else{
			fmt.Println(len10," byte written successfully...")	
		}
	}


	//Writing to file concurrently
	fmt.Println("Writing to file concurrently")
	data := make(chan int)
    done := make(chan bool)
    wg := sync.WaitGroup{}
    for i := 0; i < 10; i++ {
        wg.Add(1)
		//fmt.Println("wg= ",wg)
        go produce(data, &wg,i)
    }
    go consume(data, done)
    go func() {
		//fmt.Println("wg= ",wg)
        wg.Wait()
        close(data)
    }()
    d := <-done
    if d == true {
        fmt.Println("File written successfully")
    } else {
        fmt.Println("File writing failed")
    }

}


//Writing to file concurrently
func produce(data chan int, wg *sync.WaitGroup,i int) {  
	n := i
    data <- n
    wg.Done()
}
func consume(data chan int, done chan bool) {  
    f, err := os.Create("concurrent.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    for d := range data {
        _, err = fmt.Fprintln(f, d)
        if err != nil {
            fmt.Println(err)
            f.Close()
            done <- false
            return
        }
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        done <- false
        return
    }
    done <- true
}