package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	folderPath := os.Args[1]
	fmt.Println("Folder: ", folderPath)

	entries, err := os.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	var wg sync.WaitGroup

	ch := make(chan int)
	total := 0

	// go routines send to ch
	for _, entry := range entries {
		wg.Add(1)
		go func(e os.DirEntry) {
			defer wg.Done()
			content, err := os.ReadFile(folderPath + "/" + e.Name())
			if err != nil {
				fmt.Println("Error reading file: ", err)
				return
			}
			words := strings.Fields(string(content))
			ch <- len(words)
		}(entry)
	}

	// close channel when all goroutines finish
	go func() {
		wg.Wait()
		close(ch)
	}()

	// receive from channel until its closed
	for count := range ch {
		total += count
	}

	fmt.Println("Total words: ", total)
}
