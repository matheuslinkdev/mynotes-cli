package notes

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var ListNotes = &cobra.Command{
	Use:   "lsnt",
	Short: "List notes",
	Long:  "List all the notes that've been written",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error obtaining the user directory:", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "notes.txt")

		file, err := os.Open(fileName)
		if err != nil{
			fmt.Println("Error opening the notes file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan(){
			fmt.Println("-", scanner.Text())
		}

		if err := scanner.Err(); err != nil{
			fmt.Println("Error reading the notes file:", err)
		}
	},
}