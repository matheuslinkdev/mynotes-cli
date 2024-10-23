package cmd 

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"math/rand"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func GenerateRandomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		randomIndex := r.Intn(len(charset))
		sb.WriteByte(charset[randomIndex])
	}

	return sb.String()
}

var addBooknoteCmd = &cobra.Command{
	Use: "adb [note] [book]",
	Short: "Add a booknote",
	Long: "Add a booknote to your notes list",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error while obtaining the directory", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "booknotes.txt")
		note := args[0]
		book := args[1]
		id := GenerateRandomID(4)

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

var addNoteCmd = &cobra.Command{
	Use: "adn [phrase] [whosaid]",
	Short: "Add a note",
	Long: "Add a note to the notes file",
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error while obtaining the directory", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "notes.txt")
		phrase := args[0]
		whosaid := args[1]
		id := GenerateRandomID(4)

        file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println("Error opening the file:", err)
            return
        }
        defer file.Close()

        writer := bufio.NewWriter(file)
        _, err = writer.WriteString(fmt.Sprintf("%s - %s - %s\n", id, phrase, color.GreenString(whosaid)))
        if err != nil {
            fmt.Println("Error writing the file:", err)
            return
        }

        writer.Flush()
		time.Sleep(2 * time.Second)
        fmt.Println("Phrase added successfully!")
    },
}
