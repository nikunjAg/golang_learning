package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// embedded structs
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

const WRITE_PERMISSION = 0644

func New(Title string, Content string) (*Note, error) {
	if Title == "" || Content == "" {
		return nil, errors.New("title and content cannot be empty")
	}

	return &Note{
		Title,
		Content,
		time.Now(),
	}, nil
}

func (note Note) Save() error {

	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename)
	filename = filename + ".json"

	JSON, err := note.ToJson()

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, []byte(JSON), WRITE_PERMISSION)

	if err == nil {
		return nil
	}

	return errors.Join(errors.New("unable to write to file"), err)
}

func (note Note) ToJson() (string, error) {

	bytesArr, err := json.Marshal(note)

	if err != nil {
		return "", errors.Join(errors.New("unable to convert to json"), err)
	}

	return string(bytesArr), nil
}

func (note Note) Display() {
	fmt.Printf("Your Note -\nTitle: %v\nContent:%v\n", note.Title, note.Content)
}
