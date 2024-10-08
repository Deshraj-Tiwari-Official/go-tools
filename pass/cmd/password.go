package cmd

import (
	"fmt"
	"math/rand/v2"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a password",
  Run: generatePassword,
}

func init() {
  rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of the password")
	generateCmd.Flags().BoolP("symbols", "s", false, "Include special characters")
  generateCmd.Flags().BoolP("numbers", "n", false, "Include numbers")
}

func generatePassword(cmd *cobra.Command, args []string) {
  length, _ := cmd.Flags().GetInt("length")
  is_symbols, _ := cmd.Flags().GetBool("symbols")
  numbers, _ := cmd.Flags().GetBool("numbers")

  charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

  if is_symbols {
    charset += "!:?@#$%^&*()-_=+[]{}<>,./;"
  }
  if numbers {
    charset += "0123456789"
  }

  password := make([]byte, length)
  for i := range password {
    password[i] = charset[rand.IntN(len(charset))]
  }

  fmt.Println(string(password))
}
