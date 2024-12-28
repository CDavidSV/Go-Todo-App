package cmd

import (
	"fmt"
	"strconv"

	"github.com/CDavidSV/go-todo-app/config"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"c"},
	Args:    cobra.ExactArgs(1),
	Short:   "Mark a task as complete with the given ID",
	Long:    "Marks a task as completed with the given ID, and removes it from the list of pending tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), config.ErrorSytle.Render("Invalid task ID: %s\n"), args[0])
		}

		removed := TodoList.CompleteTask(taskID)

		if !removed {
			fmt.Fprintf(cmd.ErrOrStderr(), config.WarningStyle.Render("Task not found: %d\n"), taskID)
			return
		}

		fmt.Println(config.SuccessStyle.Render("Task marked as completed"))
	},
}
