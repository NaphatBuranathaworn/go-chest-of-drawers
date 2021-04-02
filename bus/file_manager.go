package bus

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	// "os"
	"path/filepath"
	"strings"

	"github.com/NaphatBuranathaworn/chest-of-drawers/utils"
)

var listSplitter = utils.ArrayString{"_", " ", "-", "|"}
var listExtension = utils.ArrayString{".png", ".jpg", ".jpeg", ".gif", ".xlsx", ".cpgz", ".zip", ".aab"}

type File struct {
	name string
}

// CopyFile ...;
func CopyFile(path string, folderName string) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	/* Check folder */
	abPath, err := createDirectory(path, folderName)
	
	if err != nil {
		log.Fatal(err)
	}

	/* Check by file */
	var fileSplt []string

	for _, f := range files {
		var file File
		file.name = f.Name()

		extenstion := filepath.Ext(file.name)

		if listExtension.Contains(extenstion) {

			sign, err := file.GetFileSplitter()
			if err != nil {
				log.Fatal(err)
			}

			fileSplt, err = file.SplitFileName(sign)
			if err != nil {
				log.Fatal(err)
			}

			/* TODO : หา solution ในการ แบ่ง group sub folder */
			indexForGroup, err := file.FindIndexDateInFileName(fileSplt)
			if err != nil {
				log.Fatal(err)
			}


			oldFile := filepath.Join(path, file.name)
			newFile := filepath.Join(abPath, fileSplt[indexForGroup], file.name)
			fmt.Printf("from [%s] to [%s] \n", oldFile, newFile)

			if _, err := createDirectory(abPath, fileSplt[indexForGroup]); err != nil {
				log.Fatal(err)
			}
			moveFile(oldFile, newFile)
			// fmt.Printf("from [%s] to [%s] \n", oldFile, newFile)
		
		}
	}

}

func createDirectory(path string, dir string) (string, error) {
	abPath := filepath.Join(path, dir)

	if _, err := os.Stat(abPath); !os.IsNotExist(err) {
		return abPath, err
	}

	if errCreateDir := os.MkdirAll(abPath, 0755); errCreateDir != nil {
		return abPath, errCreateDir
	}

	return abPath, nil
}

func moveFile(oldFile string, newFile string) (bool, error) {
	err := os.Rename(oldFile, newFile)
	if err != nil {
		return false, err
	}

	return true, nil
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

func (f *File) FindIndexDateInFileName(fileSplt utils.ArrayString) (int, error) {
	if f == nil {
		return 0, nil
	}

	var indexResult int

	const layoutISO = "2006-01-02"

	for i := 0; i < len(fileSplt); i++ {
		_, err := time.Parse(layoutISO, fileSplt[i])
		if err == nil {
			indexResult = i
		}
	}

	if indexResult > 0 {
		return indexResult, nil
	} else {
		return 0, nil
	}

}
