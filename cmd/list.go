package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var fileType string

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List notes or booknotes",
	Long:  "List all notes from the specified file (booknotes or general notes)",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error obtaining the user directory:", err)
			return
		}

		var fileName string
		if fileType == "booknotes" {
			fileName = filepath.Join(homeDir, "Documents", "booknotes.txt")
		} else {
			fileName = filepath.Join(homeDir, "Documents", "notes.txt")
		}

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		fmt.Println("Your notes:")

		for scanner.Scan() {
			fmt.Println("-", scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the file:", err)
		}
	},
}

func init() {
	listCmd.Flags().StringVarP(&fileType, "file", "f", "notes", "Specify the file type (booknotes or notes)")
}
