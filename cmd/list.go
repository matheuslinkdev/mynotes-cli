package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list-notes",
	Short: "List all the notes",
	Long:  "List all notes that've been addeed",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error obtaining the user directory:", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "notes.txt")

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		fmt.Println("Your book notes:")

		for scanner.Scan(){
			fmt.Println("-", scanner.Text())
		}

		if err := scanner.Err(); err != nil{
			fmt.Println("Error reading the file:", err)
		}
	},
}
