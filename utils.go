package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

var todoDir = func() string {
	str, err := os.UserHomeDir()
	if err != nil {
		panic("No home dir found")
	}
	return str + "/todo"
}
var todoDelimator = "$#||$#"

func makeDir() {
	_, err := os.Stat(todoDir())
	if err != nil {
		if os.IsNotExist(err) {
			errx := os.Mkdir(todoDir(), 0755)
			println(errx)
		}
	}
}
func GetInput() string {
	var ch string
	fmt.Print("$ ")
	inputReader := bufio.NewReader(os.Stdin)
	ch, ok := inputReader.ReadString('\n')
	if ok != nil {
		return "a"
	}
	ch = strings.Trim(ch, "\n")
	return ch

}
func Clear() {
	fmt.Print("\033c")
}
func ShowUsege() {
	color.New(color.FgGreen).Println("======================")
	fmt.Println("a: Add new TODO")
	fmt.Println("s: Show all TODO")
	fmt.Println("d: Delete TODO")
	fmt.Println("c: Complate TODO")
	fmt.Println("x: Exit")
	color.New(color.FgGreen).Println("======================")
}
func getTodoPath(f string) string {
	return todoDir() + "/" + f
}
func AddNewTodo(s string) error {
	todoFile := "todos.todo"
	makeDir()
	f, err := os.OpenFile(getTodoPath(todoFile), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(getTodoPath(todoFile))
		if err != nil {
			return err
		}
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	_, err = f.WriteString(fmt.Sprintf("\n%s %s 0", s, todoDelimator))

	return err
}
func GetTodos() []string {
	todoFile := "todos.todo"
	makeDir()
	f, err := os.Open(getTodoPath(todoFile))
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(getTodoPath(todoFile))
		if err != nil {
			fmt.Println(err)
		}
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var str string = ""
	buf := make([]byte, 1024)
	for {
		n, err1 := f.Read(buf)
		if err1 != io.EOF {
			if n > 0 && len(buf) > 0 {
				str += string(buf[:n])
				continue
			}
		}
		break
	}
	return strings.Split(str, "\n")
}
func MakeComplate(lineNum int) error {
	todos := GetTodos()
	var newTodos string = ""
	var f []string
	for idx, t := range todos {
		if idx+1 == lineNum {
			f = strings.Split(t, todoDelimator)
			newTodos = fmt.Sprintf("%s\n%s %s 1", newTodos, f[0], todoDelimator)
		} else {
			newTodos = fmt.Sprintf("%s\n%s", newTodos, t)

		}
	}
	err := saveTodos(newTodos)
	return err

}
func DeleteTodo(lineNum int) error {
	todos := GetTodos()
	var newTodos string = ""
	for idx, t := range todos {
		if idx+1 != lineNum {
			newTodos = fmt.Sprintf("%s\n%s", newTodos, t)
		}
	}
	err := saveTodos(newTodos)
	return err

}
func saveTodos(todos string) error {
	todoFile := "todos.todo"
	makeDir()
	f, err := os.OpenFile(getTodoPath(todoFile), os.O_TRUNC|os.O_WRONLY, os.ModeAppend)
	if err != nil && os.IsNotExist(err) {
		f, err = os.Create(getTodoPath(todoFile))
		if err != nil {
			return err
		}
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	_, err = f.WriteString(strings.Trim(todos, "\n"))

	return err
}
