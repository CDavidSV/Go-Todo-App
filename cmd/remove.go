package cmd

import (
	"fmt"
	"strconv"

	"github.com/CDavidSV/go-todo-app/config"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r", "rm"},
	Args:    cobra.ExactArgs(1),
	Short:   "Remove a task from your TODO list",
	Long:    "removes a task from your TODO list by providing its ID. You can get the ID by listing all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Fprintf(cmd.ErrOrStderr(), config.ErrorSytle.Render("Invalid task ID: %s\n"), args[0])
		}

		if task, deleted := TodoList.RemoveTask(taskID); deleted {
			fmt.Println(config.SuccessStyle.Render(fmt.Sprintf("Task \"%s\" (%d) removed", task.Title, task.ID)))
			return
		}

		fmt.Println(config.WarningStyle.Render(fmt.Sprintf("No tasks found matching ID: %d", taskID)))
	},
}
