package main

import (
	"chest-of-drawers/bus"
	"os"
	"bufio"
	"fmt"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var path string
    for {
        fmt.Print("Path screen shots => ");
        fmt.Scanln(&path)
        fmt.Println("Confirm path : ", path)
        fmt.Print("Please Enter to continue ...")

        scanner.Scan()
        text := scanner.Text()
        
        if text == ":wq" || len(path) > 0 {
            break
        }
    }

    // path := "/Users/naphatburanathaworn/Desktop"
    key := "Screen"
    bus.CopyFile(path, key)
}