package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/CDavidSV/go-todo-app/config"
	"github.com/CDavidSV/go-todo-app/framework"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var TodoList = framework.NewTodoList()

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Todo is a CLI task manager app with categories",
	Long:  "Todo is a complete and easy to use task management application, so you can stay organized and get things done. It allows to to create tasks and categories to organize them.",
	Run: func(cmd *cobra.Command, args []string) {
		style := lipgloss.NewStyle().
			Foreground(lipgloss.Color("99"))

		fmt.Print(style.Render(config.Logo), "\n\n\n")
		time.Sleep(100 * time.Millisecond)
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.AddCommand(completeCmd)
	rootCmd.AddCommand(listCmd)
}
