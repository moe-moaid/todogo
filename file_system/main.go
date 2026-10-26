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

func editListItem(f string, oldString, newString string) {
	data, err := os.ReadFile(f)
	if (err != nil) {
		fmt.Println("error wile reading file:: ", err)
	}
	isStringExist := strings.Contains(string(data), oldString)
	if isStringExist {
		fmt.Println("isStringExist:: ", isStringExist)
		fmt.Println("old Data: ", string(data))
		replacedData := strings.Replace(string(data), oldString, newString, 1)
		fmt.Println("replaced Data: ", replacedData)
		os.WriteFile("text.txt", []byte(replacedData), 0644)	
	} else {
		fmt.Println("no match was found bruh!")
	}
}

func removeItem (item string, f string) {
	data, _ := os.ReadFile(f)
}

func main() {
	file,_ := os.OpenFile("text.txt", O_APPEND|O_CREATE|O_RDWR, 0644)
	
	saveListItem(file)
	editListItem("text.txt", "- just a new style", "- new str")
	
	defer file.Close()
}
