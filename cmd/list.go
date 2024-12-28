package cmd

import (
	"fmt"

	"github.com/CDavidSV/go-todo-app/config"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List all tasks",
	Long:    "Lists all tasks, optionally filtering by category or completion status.",
	Run: func(cmd *cobra.Command, args []string) {
		category, _ := cmd.Flags().GetString("category")
		completed, _ := cmd.Flags().GetBool("completed")

		var tasks [][]string
		if category == "" {
			tasks = TodoList.ListAllTasks(completed)
		} else {
			tasks = TodoList.ListTasks(category, completed)
		}

		if len(tasks) == 0 {
			fmt.Println(config.InfoStyle.Render("No tasks found"))
			return
		}

		t := table.New().
			Border(lipgloss.RoundedBorder()).
			BorderStyle(config.TableBorderStyle).
			StyleFunc(func(row, col int) lipgloss.Style {
				switch {
				case row == -1:
					return config.TableHeaderStyle
				case row%2 == 0 && (col == 5 || col == 0):
					return config.TableEvenRowStyle.Align(lipgloss.Center)
				case row%2 != 0 && (col == 5 || col == 0):
					return config.TableOddRowStyle.Align(lipgloss.Center)
				case row%2 == 0:
					return config.TableEvenRowStyle
				default:
					return config.TableOddRowStyle
				}
			}).
			Headers("ID", "TITLE", "DESCRIPTION", "CREATED", "CATEGORY", "COMPLETED").
			Rows(tasks...).
			Width(100)

		fmt.Println(t)
	},
}

func init() {
	listCmd.Flags().BoolP("completed", "x", false, "Show completed tasks")
	listCmd.Flags().StringP("category", "c", "", "Filter by category")
}
