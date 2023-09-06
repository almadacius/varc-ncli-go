package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	filePath := "myfile.txt"

	// Open the file
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Request an exclusive lock on the file
	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX)
	if err != nil {
		fmt.Println("Error locking file:", err)
		return
	}
	defer syscall.Flock(int(file.Fd()), syscall.LOCK_UN)

	// Perform the write operation
	_, writeErr := file.WriteString("Hello, World!\n")
	if writeErr != nil {
		fmt.Println("Error writing to file:", writeErr)
		return
	}
	fmt.Println("Write operation completed.")
}
