package main

import (
	"fmt"
	"os"
	"syscall"
)

func saveListItem() {
	//	os.WriteFile("todos.txt", []byte ("new task added"),0666)
	const (
		O_APPEND int = syscall.O_APPEND
		O_CREATE int = syscall.O_CREAT
		O_WRONLY int = syscall.O_WRONLY
	)
	f, _ := os.OpenFile("todo.txt", O_APPEND|O_CREATE|O_WRONLY, 0644)
	_, err := f.Write([]byte("- new task added to the file shitze \n"))
	
	if (err != nil) {
		fmt.Println("error while writing your file:: ", err)
	}
	f.Close()
	fmt.Println("added new element to file")
}

func main() {
	saveListItem()
}
