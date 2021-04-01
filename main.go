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

    for {
        fmt.Print("Press Enter to execute... ")
        scanner.Scan()
        text := scanner.Text()
        if text == ":wq" || len(text) == 0 {
            break
        }
    }

    // path, err := os.Getwd()
    // if err != nil {
    //     log.Println(err)
    // }

    path := "/Users/naphatburanathaworn/Desktop"
    key := "Screen"
    bus.CopyFile(path, key)
}