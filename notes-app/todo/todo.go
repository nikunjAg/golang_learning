package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

// embedded structs
type Todo struct {
	Text string `json:"text"`
}

const WRITE_PERMISSION = 0644

func New(Text string) (*Todo, error) {
	if Text == "" {
		return nil, errors.New("text cannot be empty")
	}

	return &Todo{
		Text,
	}, nil
}

func (todo Todo) Save() error {

	filename := strings.ReplaceAll(todo.Text, " ", "_")
	filename = strings.ToLower(filename)
	filename = filename + ".json"

	JSON, err := todo.ToJson()

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, []byte(JSON), WRITE_PERMISSION)

	if err == nil {
		return nil
	}

	return errors.Join(errors.New("unable to write to file"), err)
}

func (todo Todo) ToJson() (string, error) {

	bytesArr, err := json.Marshal(todo)

	if err != nil {
		return "", errors.Join(errors.New("unable to convert to json"), err)
	}

	return string(bytesArr), nil
}

func (todo Todo) Display() {
	fmt.Printf("Your Todo -\nText: %v\n", todo.Text)
}
