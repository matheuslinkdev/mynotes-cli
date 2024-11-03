package handlers

import (
	"github.com/matheuslinkdev/mynotes-cli/handlers/bookill"
	"github.com/matheuslinkdev/mynotes-cli/handlers/books"
	"github.com/matheuslinkdev/mynotes-cli/handlers/notes"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mynotes",
	Short: "CLI to write down book quotes and citations.",
	Long:  "My CLI for adding and listing quotes from books or people.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(books.AddBooknote)
	rootCmd.AddCommand(books.RemoveBooknote)
	rootCmd.AddCommand(books.ListBooks)
	rootCmd.AddCommand(notes.AddNote)
	rootCmd.AddCommand(notes.RemoveNote)
	rootCmd.AddCommand(notes.ListNotes)
	rootCmd.AddCommand(bookill.AddBookill)
	rootCmd.AddCommand(bookill.RemoveBookill)
	rootCmd.AddCommand(bookill.ListBookill)
}
