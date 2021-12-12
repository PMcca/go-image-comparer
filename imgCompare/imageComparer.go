package imgCompare

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
)

func CompareImagesFromDirs(dir1 string, dir2 string) error {
	dir1Images, err := loadImages(dir1)
	if err != nil {
		return err
	}

	dir2Images, err := loadImages(dir2)
	if err != nil {
		return err
	}

	fmt.Println(dir1Images)
	fmt.Println(dir2Images)

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
