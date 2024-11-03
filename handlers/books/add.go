package books

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

var AddBooknote = &cobra.Command{
	Use:   "adbn [note] [book]",
	Short: "Add a booknote",
	Long:  "Add a booknote to notes list",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error while obtaining the directory", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "booknotes.txt")
		note := args[0]
		book := args[1]
		id := utils.GenerateRandomID(4)

		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		_, err = writer.WriteString(fmt.Sprintf("%s - %s - %s\n", id, note, color.YellowString(book)))
		if err != nil {
			fmt.Println("Error writing the file:", err)
			return
		}

		writer.Flush()
		time.Sleep(2 * time.Second)
		fmt.Println("Booknote added successfully!")
	},
}
