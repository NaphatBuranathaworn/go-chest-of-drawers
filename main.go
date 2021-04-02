package main

import (
	"bufio"

	"fmt"
	// "log"
	"os"

	"github.com/NaphatBuranathaworn/chest-of-drawers/bus"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var folderName string
    for {
        // fmt.Print("Press Enter to execute... ")
        // scanner.Scan()
        // text := scanner.Text()
        // if text == ":wq" || len(text) == 0 {
        //     break
        // }

        fmt.Print("Folder name => ")
		fmt.Scanln(&folderName)

		fmt.Printf("Confirm folder name : %s\n", folderName)
		fmt.Print("Please Enter to continue ...")

		scanner.Scan()
		text := scanner.Text()

		if text == ":wq" || len(folderName) > 0 {
			break
		}
    }

    path, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
    }

    // path := "/Users/naphatburanathaworn/Desktop"
    // key := "Screen"
    bus.CopyFile(path, folderName)
}