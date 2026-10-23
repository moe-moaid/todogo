package main

import (
	"fmt"
	"os"
	"syscall"
	// "io"
	"strings"
)

const (
	O_APPEND int = syscall.O_APPEND
	O_CREATE int = syscall.O_CREAT
	O_RDWR int = syscall.O_RDWR
)

func saveListItem(f *os.File) {
	_, err := f.Write([]byte("- new task added to the file shitze \n"))
	
	if (err != nil) {
		fmt.Println("error while writing your file:: ", err)
	}
	
	fmt.Println("added new element to file")
}

func editListItem(f string) bool{
	data, err := os.ReadFile(f)
	// r := strings.NewReader("new string")
	// data, err := io.ReadAll(r)
	if (err != nil) {
		fmt.Println("error wile reading file:: ", err)
	}
	replacedData := strings.Replace(string(data), "- new task added to the file shitze", "- just a new style", -1)
	fmt.Printf("data extracted from file:: %s \n", data)
	fmt.Printf("replaced Data:: %s \n", replacedData)
	os.WriteFile("text.txt", []byte(replacedData), 0644)
	fmt.Printf("new file content:: %s", []byte(replacedData))
	return strings.Contains(string(data), "task")
}

func main() {
	file,_ := os.OpenFile("text.txt", O_APPEND|O_CREATE|O_RDWR, 0644)
	
	saveListItem(file)
	editListItem("text.txt")
	
	defer file.Close()
}
