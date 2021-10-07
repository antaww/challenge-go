package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input_files := os.Args[1:]
	if len(input_files) == 0 {
		fmt.Println("No file")
	} else {
		fmt.Println(input_files[0])
		content, err := ioutil.ReadFile(input_files[0])
		if err != nil {
			fmt.Println("Can't read ", input_files[0], "Error : ", err)
		} else {
			fmt.Println("File", input_files[0], ": \n", string(content))
		}
	}
}
