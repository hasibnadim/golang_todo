package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func AddNew() {
	Clear()
	color.New(color.FgBlue).Println("Write New TODO:")
	todo := GetInput()
	err := AddNewTodo(todo)
	if err != nil {
		fmt.Println(err)
		color.New(color.FgRed).Println("Todo insert failed!")
	} else {
		color.New(color.FgGreen).Println("Todo insert successful.")
	}

	ShowUsege()
}
func ShowAll() {
	todos := GetTodos()
	Clear()
	var f []string
	color.New(color.BgBlue).Add(color.FgGreen).Println("P=Pending | C=Complate")
	fmt.Println("_________________________________")
	for _, t := range todos {
		f = strings.Split(t, "$#||$#")
		if len(f) == 2 {
			i, _ := strconv.Atoi(strings.Trim(f[1], " "))
			if i == 1 {
				fmt.Printf("C| %s\n", f[0])
			} else {
				fmt.Printf("P| %s\n", f[0])
			}
		}
	}
	fmt.Println("_________________________________")
	ShowUsege()
}
func ComplateTodo() {
EXIT:
	for {

		todos := GetTodos()
		Clear()
		var f []string
		ix := 1
		var ixi = make(map[int]int)
		fmt.Println("Let's Complate")
		fmt.Println("_________________________________")
		fmt.Println("")
		for idx, t := range todos {
			f = strings.Split(t, "$#||$#")
			if len(f) == 2 {
				i, _ := strconv.Atoi(strings.Trim(f[1], " "))
				if i == 0 {
					ixi[ix] = idx + 1
					fmt.Printf("%d| %s\n", ix, f[0])
					ix++
				}
			}
		}

		fmt.Println("_________________________________")
		color.New(color.FgBlue).Println("Enter selected todo number (x for exit):")
		selected := GetInput()
		i, err := strconv.Atoi(selected)
		if err != nil {
			Clear()
			ShowUsege()
			break EXIT
		}

		err = MakeComplate(ixi[i])
		if err != nil {
			fmt.Println(err)
			break EXIT
		}

	}
}
func DeleteTodoView() {
EXIT:
	for {

		todos := GetTodos()
		Clear()
		var f []string
		fmt.Println("Let's Delete")
		fmt.Println("_________________________________")
		fmt.Println("")
		for idx, t := range todos {
			f = strings.Split(t, "$#||$#")
			if len(f) == 2 {
				fmt.Printf("%d| %s\n", idx+1, f[0])
			}
		}

		fmt.Println("_________________________________")
		color.New(color.FgBlue).Println("Enter selected todo number (x for exit):")
		selected := GetInput()
		i, err := strconv.Atoi(selected)
		if err != nil {
			Clear()
			ShowUsege()
			break EXIT
		}

		err = DeleteTodo(i)
		if err != nil {
			fmt.Println(err)
			break EXIT
		}

	}
}
