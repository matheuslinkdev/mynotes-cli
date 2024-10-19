package cmd

import (
    "bufio"
    "fmt"
    "os"
    "path/filepath"
    "strings"

    "github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
    Use:   "remove [phrase]",
    Short: "Remove uma frase de livro",
    Long:  `Remove uma frase do arquivo de anotações.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        homeDir, err := os.UserHomeDir()
        if err != nil {
            fmt.Println("Erro ao obter o diretório do usuário:", err)
            return
        }

        fileName := filepath.Join(homeDir, "Documents", "notes.txt")

        file, err := os.Open(fileName)
        if err != nil {
            fmt.Println("Erro ao abrir o arquivo:", err)
            return
        }
        defer file.Close()

        var phrases []string
        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            phrases = append(phrases, scanner.Text())
        }
        if err := scanner.Err(); err != nil {
            fmt.Println("Erro ao ler o arquivo:", err)
            return
        }

        phraseToRemove := args[0]
        found := false

        for i, phrase := range phrases {
            if strings.HasPrefix(phrase, phraseToRemove) {
                phrases = append(phrases[:i], phrases[i+1:]...)
                fmt.Println("Frase removida com sucesso!")
                found = true
                break
            }
        }

        if !found {
            fmt.Println("Frase não encontrada.")
            return
        }

        file, err = os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY, 0644)
        if err != nil {
            fmt.Println("Erro ao abrir o arquivo:", err)
            return
        }
        defer file.Close()

        writer := bufio.NewWriter(file)
        for _, phrase := range phrases {
            _, err := writer.WriteString(phrase + "\n")
            if err != nil {
                fmt.Println("Erro ao escrever no arquivo:", err)
                return
            }
        }
        writer.Flush()
    },
}

func init() {
    rootCmd.AddCommand(removeCmd)
}
