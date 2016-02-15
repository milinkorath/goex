package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println(os.Args[1:])
	fmt.Println(os.ModeDevice)

	file, err := os.Open("data/hello.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = file.Stat()
	f1 := os.Stdin

	info, _ := f1.Stat()
	fmt.Println(f1.Name(), info.Mode())
	fmt.Println(os.ModeDevice & info.Mode())

	inputFile := os.Stdin
	// fmt.Println(file.Name())
	// fileMode, _ := file.Stat()
	// fmt.Println(fileMode.Mode())

}
