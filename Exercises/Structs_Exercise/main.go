package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData() (string, string, error) {
	noteTitle, err := getUserInput("Note title:")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	noteContent, err := getUserInput("Note content")
	if err != nil {
		fmt.Println(err)
		return "", "", err
	}
	return noteTitle, noteContent, nil
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	var userInputValue string
	fmt.Scanln(&userInputValue)
	if userInputValue == "" {
		return "", errors.New("Invalid input")
	}
	return userInputValue, nil
}
