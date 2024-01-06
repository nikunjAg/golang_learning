package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

// Embedded interfaces
type outputable interface {
	Display()
	saver
}

func readStringFromStdin(valuePtr *string) {
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		*valuePtr = ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	*valuePtr = text
}

func fetchNoteData() (string, string) {
	var title, content string

	fmt.Printf("Please enter note title: ")
	readStringFromStdin(&title)

	fmt.Printf("Please enter note content: ")
	readStringFromStdin(&content)

	return title, content
}

func fetchTodoData() string {
	var text string

	fmt.Printf("Please enter todo text: ")
	readStringFromStdin(&text)

	return text
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed saving the data!")
		return err
	}

	fmt.Println("Saved data successfully")
	return nil
}

func outputData(data outputable) {
	err := saveData(data)

	if err != nil {
		fmt.Println(err)
		return
	}
	data.Display()

}

func main() {

	title, content := fetchNoteData()
	var created_note *note.Note
	created_note, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	text := fetchTodoData()
	var created_todo *todo.Todo
	created_todo, err = todo.New(text)

	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(created_note)
	outputData(created_todo)
}
