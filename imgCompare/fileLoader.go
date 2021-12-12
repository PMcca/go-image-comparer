package imgCompare

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
)

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
