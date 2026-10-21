package main

import (
	"fmt"
	"os"
	"io"
)

func saveListItem() {
	//	os.WriteFile("todos.txt", []byte ("new task added"),0666)
	f, _ := os.Open("todo.txt")
	_, err := io.WriteString("todo.txt", "new task added to the file shitze \n")
	if (err != nil) {
		fmt.Println("error while writing your file:: ", err)
	}
	f.Close()
	fmt.Println("added new element to file")
}

func main() {
	saveListItem()
}
