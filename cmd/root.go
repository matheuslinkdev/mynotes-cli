package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "mybooknotes",
	Short: "CLI to write down sentences of books",
	Long: "My CLI to add and list phrases and citations of books",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil{
		panic(err)
	}
}

func init(){
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(listCmd)
}