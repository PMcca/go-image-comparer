package imgCompare

import (
	"fmt"
	"image"
	_"image/jpeg"
	_"image/png"
	"log"
	"os"
)

var dir1ImgHashes = make(map[string]bool)

func CompareImagesFromDirs(dir1 string, dir2 string) {
	dir1Contents := getDirFilenames(dir1)
	//dirContents2 := getDirFilenames(dir2)

	loadDir1Hashes(dir1Contents)
	//openImage(dir1Contents[0])
	//fmt.Println(dirContents2)

}

// Reads dir and returns list of full filename paths
func getDirFilenames(dir string)[]string {
	fmt.Printf("Reading directory %s\n", dir)

	filenames := make([]string, 0)
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		filenames = append(filenames, dir + "/" + file.Name())
	}
	return filenames
}

// Computes a hash for each image in dir1 and puts hash into dir1ImgHashes
func loadDir1Hashes(dir1Filenames []string) {
	for _, path := range dir1Filenames {
		fmt.Println("Opening and decoding image " + path)
		// Open file stream
		file, err := os.Open(path)
		if err != nil {
			printAndExit("ERROR reading file " + path, err)
		}

		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				printAndExit("ERROR closing file " + file.Name(), err)
			}
		}(file)

		// Decode image
		img, _, err := image.Decode(file)
		if err != nil {
			printAndExit("ERROR decoding image " + file.Name(), err)
		}

		width, height := img.Bounds().Max.X, img.Bounds().Max.Y
		//fmt.Println(width)
		//fmt.Println(height)
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				fmt.Println(img.At(x, y))
			}
		}
	}
}

func printAndExit(msg string, err error) {
	fmt.Println(msg, err)
	os.Exit(1)
}
