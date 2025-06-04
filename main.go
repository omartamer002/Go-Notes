// This is a demo note app that prompts the user to enter their note title and content.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
)

func getNoteData() (string, string) {
	// Prompt the user to enter their note title
	title := getUserInput("Enter Note Title: ")
	// Prompt the user to enter their note content
	content := getUserInput("Enter Note Content: ")

	return title, content
}

func main() {
	title, content := getNoteData()
	userNote, err := note.NewNote(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.DisplayNote()
	err = userNote.SaveNote()
	if err != nil {
		fmt.Println("Failed to save the note!")
		return
	}
	fmt.Println("Note saved successfully!")
}

// Creating a helper function to get user input
func getUserInput(prompt string) string {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
