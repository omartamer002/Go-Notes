// Creating a new package named "note" that will store note struct
package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// creating a struct to store note data
type Note struct {
	Title     string
	Content   string
	CreatedAt time.Time
}

// creating a function to display a user's note
func (note Note) DisplayNote() {
	fmt.Printf("Your note titled %v has the following content: \n\n%v\n\n", note.Title, note.Content)
}

// Saving the note content into a file (JSON file)
func (note Note) SaveNote() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}

// creating a constructor function to initialize the note struct
func NewNote(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input")
	}
	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
