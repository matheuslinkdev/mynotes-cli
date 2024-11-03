package bookill

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/fatih/color"
	"github.com/matheuslinkdev/mynotes-cli/utils"
	"github.com/spf13/cobra"
)

var AddBookill = &cobra.Command{
	Use:   "adbk [book] [author]",
	Short: "Add a bookill",
	Long:  "Add a finished book to the list",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error while obtaining the directory", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "bookill.txt")
		book := args[0]
		author := args[1]
		id := utils.GenerateRandomID(3)

		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(fmt.Sprintf("%s - %s - %s\n", id, color.GreenString(book), author))
		if err != nil {
			fmt.Println("Error writing the file:", err)
			return
		}

		writer.Flush()
		time.Sleep(2 * time.Second)
		fmt.Println("Bookill added successfully!")
	},
}
