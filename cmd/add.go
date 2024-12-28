package cmd

import (
	"fmt"

	"github.com/CDavidSV/go-todo-app/config"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new task to your TODO list",
	Long:    "Adds a new task to your TODO list. You can specify a category for the task.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		description, _ := cmd.Flags().GetString("description")
		category, _ := cmd.Flags().GetString("category")

		if description == "" {
			description = "No description"
		}

		defer fmt.Println(config.SuccessStyle.Render("Task added successfully"))
		if category == "" {
			TodoList.AddGenericTask(name, description)
			return
		}

		TodoList.AddTask(name, description, category)
	},
}

func init() {
	addCmd.Flags().StringP("name", "n", "", "Name of the task")
	addCmd.Flags().StringP("description", "d", "", "Description of the task")
	addCmd.Flags().StringP("category", "c", "", "Name of the category to which the task belongs")

	addCmd.MarkFlagRequired("name")
}
