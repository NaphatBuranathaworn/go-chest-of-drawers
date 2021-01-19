package bus

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// CopyFile ...;
func CopyFile(path string, pattern string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		var txtArr ArrayString
		txtArr = strings.Fields(f.Name())

		if txtArr.contains(pattern) {
			abPath := filepath.Join(path, "Screen Shots", txtArr[2])

			if err = os.MkdirAll(abPath, 0755); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("create dir : %s success. \n", abPath)

			oldFile := filepath.Join(path, f.Name())
			newFile := filepath.Join(abPath, f.Name())

			fmt.Printf("move file [%s] to [%s] success.\n", oldFile, newFile)

			err := os.Rename(oldFile, newFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}

// ArrayString ...
type ArrayString []string

func (s ArrayString) contains(str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
