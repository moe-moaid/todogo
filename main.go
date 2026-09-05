package main

import (
	"fmt"
	"bufio"
	"os"
)

type choices struct {
	num int
	action string
}

func printListItems(someList[]string) {
	for i := 0; i < len(someList); i++ {
		fmt.Printf("- %s", someList[i])
	}
}


func main() {
	instructions := [3]choices {
		choices {num: 1, action: "Add a task"},
		choices {num: 2, action: "Remove a task"},
		choices {num: 3, action:"List tasks"},
	}
	var defaultTasks []string
	var choiceInput int
	var taskToRemove int
	if len(defaultTasks) == 0 {
		fmt.Println("You don't have tasks currently!")
	} else {
	fmt.Println("Current Tasks:")	
	}

	for choiceInput != 3 {
		fmt.Println("choose what you wanna do next: ")
		if len(defaultTasks) == 0 {
			fmt.Printf("> %d %s \n", instructions[0].num, instructions[0].action)
			fmt.Scan(&choiceInput)
			fmt.Println("Enter a starter task to your list: ")
			scannedInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				fmt.Println("fuckin shit")
			} else {
				fmt.Println(scannedInput)
			}
			newArr := append(defaultTasks, scannedInput)
			fmt.Println("Task added to your list: ")
			defaultTasks = newArr
			fmt.Printf("- %s ", defaultTasks[0])
		} else {
			for i := 0; i < len(instructions); i++ {
				fmt.Printf("> %d %s \n", instructions[i].num, instructions[i].action)
			}
			fmt.Scan(&choiceInput)
			if choiceInput == 1 {
				fmt.Println("nice, Add one more task bellow: ")
				line, err := bufio.NewReader(os.Stdin).ReadString('\n')
				if err != nil {
					fmt.Println(err)
				}
				newArr := append(defaultTasks, line)
				defaultTasks = newArr
				fmt.Println("Task Added successfully! \n your updated list is:")
				printListItems(defaultTasks)
				
			} else if choiceInput == 2 {
				fmt.Println("Enter the number of the task that you want to remove")
				fmt.Scan(&taskToRemove)
				newArr := append(defaultTasks[:taskToRemove], defaultTasks[taskToRemove + 1:]...)
				fmt.Println("Task removed Successfully! \n your update list is: ")
				defaultTasks = newArr
				printListItems(defaultTasks)
			} else if choiceInput ==3 {
				fmt.Println("here is your fucking list")
				printListItems(defaultTasks)
			}	
		}
	}
}
