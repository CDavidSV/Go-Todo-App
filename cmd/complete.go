package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CDavidSV/go-todo-app/config"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:     "complete",
	Aliases: []string{"c"},
	Args:    cobra.MinimumNArgs(1),
	Short:   "Mark tasks as complete with the given ID or series of IDs",
	Long:    "Allows you to marks a task or multiple tasks as complete by providing their IDs. You can get the ID by listing all tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		var tasksNotFound []string
		var completedTasks []string
		for _, arg := range args {
			taskID, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), config.ErrorStyle.Render("Invalid task ID: %s\n"), arg)
				return
			}

			removed := TodoList.SetTaskCompletion(taskID, true)

			if !removed {
				tasksNotFound = append(tasksNotFound, arg)
			} else {
				completedTasks = append(completedTasks, arg)
			}
		}

		if len(tasksNotFound) > 0 {
			fmt.Println(config.WarningStyle.Render(fmt.Sprintf("No tasks found matching ID(s): %s", strings.Join(tasksNotFound, ", "))))
		}

		if len(tasksNotFound) != len(args) {
			fmt.Println(config.SuccessStyle.Render(fmt.Sprintf("Task(s) %s marked as completed", strings.Join(completedTasks, ", "))))
		}
	},
}

func init() {
}
