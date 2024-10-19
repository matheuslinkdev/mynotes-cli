package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add [phrase] [book]",
	Short: "Add a note",
	Long: "Add a note to the notes file",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()

		if err != nil{
			fmt.Println("Error while obtaining the directory", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "notes.txt")

		 phrase := args[0]
		 book := args[1]

        file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println("Error opening the file:", err)
            return
        }
        defer file.Close()

        writer := bufio.NewWriter(file)
        _, err = writer.WriteString(phrase + " - " + color.YellowString(book))
        if err != nil {
            fmt.Println("Error writing the file:", err)
            return
        }

        writer.Flush()
		time.Sleep(2 *time.Second)
        fmt.Println("Phrase added succesfully!")
    },
}