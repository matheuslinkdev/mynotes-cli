package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RemoveNote(fileName, noteToRemove string) error {
	file, err := os.Open(fileName)

	if err != nil{
		return fmt.Errorf("error opening the file: %v", err)
	}

	defer file.Close()

	var notes []string

	scanner := bufio.NewScanner(file)

	 for scanner.Scan() {
        notes = append(notes, scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return fmt.Errorf("error while reading file: %v", err)
    }

    found := false
    for i, phrase := range notes {
        if strings.HasPrefix(phrase, noteToRemove) {
            notes = append(notes[:i], notes[i+1:]...)
            fmt.Println("Phrase removed successfully!")
            found = true
            break
        }
    }

    if !found {
        return fmt.Errorf("phrase not found")
    }

    file, err = os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("error opening the file for writing: %v", err)
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    for _, phrase := range notes {
        _, err := writer.WriteString(phrase + "\n")
        if err != nil {
            return fmt.Errorf("error writing to the file: %v", err)
        }
    }
    writer.Flush()

    return nil

}