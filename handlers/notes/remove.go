package notes

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matheuslinkdev/mynotes-cli/utils"
	"github.com/spf13/cobra"
)

var RemoveNote = &cobra.Command{
	Use:   "rmn [note]",
	Short: "Removes a note",
	Long:  `Removes a note from the notes file.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error while obtaining the user folder:", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "notes.txt")
		if err := utils.RemoveNote(fileName, args[0]); err != nil {
			fmt.Println(err)
		}
	},
}
