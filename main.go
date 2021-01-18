package main

import (
	"chest-of-drawers/bus"
	"os"
	"bufio"
	"fmt"
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

    path := "/Users/naphatburanathaworn/Desktop"
    key := "Screen"
    bus.CopyFile(path, key)
}