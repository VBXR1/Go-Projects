package main

import (
	"fmt"
	"time"
	"image"
	"os"
	"syscall"
	"unsafe"
)

func main() {
	func load(filename string) (image.Image, error) {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		img, _, err := image.Decode(file)
		return img, err
	}

	for i := 0; i < 10000000; i++ {

		fmt.Printf("%v\n", i)
		if i == 5000000 {
			time.Sleep(2 * time.Second)
			print("B do Q psh do Q\n")
			time.Sleep(2 * time.Second)
		}

	}
}
