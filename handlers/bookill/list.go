package bookill

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var ListBookill = &cobra.Command{
	Use:   "lsbk",
	Short: "List bookill",
	Long:  "List all the finished books",
	Run: func(cmd *cobra.Command, args []string) {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error obtaining the user directory:", err)
			return
		}

		fileName := filepath.Join(homeDir, "Documents", "bookill.txt")

		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening the booknotes file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		counter := 0

		for scanner.Scan() {
			fmt.Println("-", scanner.Text())
			counter++
		}

		fmt.Println("Number of finished books:", color.GreenString(strconv.Itoa(counter)))

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading the bookill file:", err)
		}
	},
}
