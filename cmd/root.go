package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "mynotes",
	Short: "CLI to write down book quotes and citations.",
	Long: "My CLI for adding and listing quotes from books or people.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil{
		panic(err)
	}
}

func init(){
	rootCmd.AddCommand(addBooknoteCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addNoteCmd)
	rootCmd.AddCommand(removeBooknoteCmd)
    rootCmd.AddCommand(removeNoteCmd)
}