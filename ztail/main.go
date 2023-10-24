package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(os.Args[0])
	} else {
		num, err := strconv.Atoi(os.Args[1])

		if err != nil {
			fmt.Println(err.Error())
		} else {
			if len(os.Args[2:]) == 0 {
				fmt.Println("File name missing")
			} else {
				for _, res := range os.Args[2:] {
					ztail(res, num)
				}
			}
		}
	}
}

func ztail(s string, numBytes int) {
	file, err := os.Open(s)
	if err != nil {
		fmt.Println(err.Error())
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s\n", data[len(data)-numBytes:])
}
