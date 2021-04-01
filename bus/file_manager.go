package bus

import (
	"fmt"
	"io/ioutil"
	"log"

	// "os"
	"path/filepath"
	"strings"

	"github.com/NaphatBuranathaworn/chest-of-drawers/utils"
)

// type ArrayString []string

var listSplitter = utils.ArrayString{"_", " ", "-", "|"}
var listExtension = utils.ArrayString{".png", ".jpg", ".jpeg", ".gif"}

type File struct {
	name        string
	currentPath string
	targetPath  string
}

// CopyFile ...;
func CopyFile(path string, pattern string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	var fileSplt []string

	for _, f := range files {
		var file File
		file.name = f.Name()

		extenstion := filepath.Ext(file.name)

		if listExtension.Contains(extenstion) {

			sign, err := file.GetFileSplitter();
			if err != nil {
				log.Fatal(err)
			}

			fileSplt, err = file.SplitFileName(sign)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("test : ", fileSplt)

		}

	
	}

}

func (f *File) GetFileSplitter() (string, error) {
	if f == nil {
		return " ", nil
	}

	for _, sign := range listSplitter {
		if strings.Contains(f.name, sign) {
			return sign, nil
		}
	}

	return " ", nil
}

func (f *File) SplitFileName(key string) (utils.ArrayString, error) {
	if f == nil {
		return []string{}, nil
	}

	fileArr := strings.Split(f.name, key)
	return fileArr, nil
}



