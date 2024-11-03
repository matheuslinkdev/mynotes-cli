package bookill

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/matheuslinkdev/mynotes-cli/utils"
	"github.com/spf13/cobra"
)

var RemoveBookill = &cobra.Command{
    Use:   "rmbk [id]",
    Short: "Removes a bookill",
    Long:  `Removes a finished book from the list.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            fmt.Println("Error while obtaining the user folder:", err)
            return
        }

        fileName := filepath.Join(homeDir, "Documents", "bookill.txt")
        if err := utils.RemoveNote(fileName, args[0]); err != nil {
            fmt.Println(err)
        }
    },
}