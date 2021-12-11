package imgCompare

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"time"
)

func CompareImagesFromDirs(dir1 string, dir2 string) error {
	dir1Images, err := loadImages(dir1)
	if err != nil {
		return err
	}

	fmt.Println(dir1Images)

	return nil

}

//
//// Computes a hash for each image in dir1 and puts hash into dir1ImgHashes
//func loadDir1Hashes(dir1Filenames []string) {
//		gray := toGrayscale(img)
//
//		f, err := os.Create("some_img.png")
//		if err != nil {
//			printAndExit("Could not create file", err)
//		}
//
//		defer f.Close()
//		if err := png.Encode(f, gray); err != nil {
//			printAndExit("Could not save image ", err)
//		}
//	}
//}

// loadImages returns all images recursively found in the given dir
func loadImages(dir string) ([]image.Image, error) {
	started := time.Now()
	defer func(s time.Time) {
		fmt.Printf("took %v", time.Now().Sub(s))
	}(started)

	var images []image.Image
	dirQueue := []string{dir}

	log.Printf("Reading contents of %s\n", dir)

	// while there are directories to read
	for len(dirQueue) > 0 {
		d := dirQueue[0]
		dirQueue = dirQueue[1:] // Pop from queue

		dirContents, err := os.ReadDir(d)
		if err != nil {
			return nil, fmt.Errorf("error reading dir %s, %w", d, err)
		}

		for _, f := range dirContents {
			// Get absolute path of file found
			path, err := filepath.Abs(d + "/" + f.Name())
			if err != nil {
				return nil, fmt.Errorf("could not create path for file %s, %w", f.Name(), err)
			}

			if f.IsDir() {
				dirQueue = append(dirQueue, path) // Add dir to queue
			} else {
				img, err := openImage(path)
				if err != nil {
					return nil, fmt.Errorf("could not open file %s, %w", f.Name(), err)
				}

				images = append(images, img)
			}
		}
	}

	return images, nil
}

func openImage(file string) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Println(err)
		}
	}(f)

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("could not decode file %s to image, %w", file, err)
	}

	return img, nil
}
