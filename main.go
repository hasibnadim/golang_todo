package main

import "fmt"

func main() {
	var ch string
	Clear()
	ShowUsege()
EXIT:
	for {
		ch = GetInput()
		switch ch {
		case "a":
			AddNew()
		case "s":
			ShowAll()
		case "d":
			DeleteTodoView()
		case "c":
			ComplateTodo()
		case "x":
			break EXIT
		default:
			ShowUsege()
		}
	}
	fmt.Println("Bye Bye!")

}
