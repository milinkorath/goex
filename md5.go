package main

import (
	"crypto/md5"
	"fmt"
	"os"
)

func main() {
	info, _ := os.Stdin.Stat()
	fmt.Println(info.Mode())
	fmt.Printf("%s %T\n", os.ModeCharDevice, os.ModeCharDevice)
	fmt.Println(info.Mode())
	if (info.Mode() & os.ModeCharDevice) == os.ModeCharDevice {
		fmt.Println("******************")
	}
	showId()
	showId1()
}

func showId() {
	hash := md5.New()
	hash.Write([]byte("hello-world"))
	md := hash.Sum(nil)
	checksum := fmt.Sprintf("%x", md)
	fmt.Printf("%x", md)
	fmt.Println(len(checksum))
}

func showId1() {

	checksum := md5.Sum([]byte("hello-world"))
	fmt.Println(checksum)
	fmt.Printf("%x\n", checksum)
	fmt.Println(len(checksum))
}
