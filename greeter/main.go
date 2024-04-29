package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{
		Use:   "greeter",
		Short: "A basic CLI example",
		Long:  "A basic CLI example using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Greeter!")
		},
	}

	greetCmd := &cobra.Command{
		Use:   "greet",
		Short: "A greet command",
		Long:  "Greet someone by their name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("hello, %s!\n", args[0])
		},
	}

	rootCmd.AddCommand(greetCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
