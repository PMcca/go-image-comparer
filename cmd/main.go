package main

import (
	"go-image-comparer/imgCompare"
	"log"
)

func main() {
	err := imgCompare.CompareImagesFromDirs("F:\\Pictures\\gotst\\dir1", "F:\\Pictures\\gotst\\dir2")
	if err != nil {
		log.Fatalf("%v", err)
	}
}
