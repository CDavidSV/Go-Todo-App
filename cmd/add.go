package cmd

import (
	"fmt"
	"log"

	"github.com/CDavidSV/go-todo-app/config"
	"github.com/CDavidSV/go-todo-app/ui"
	tea "github.com/charmbracelet/bubbletea"
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

		var output ui.TextInputValue
		if name == "" {
			p := tea.NewProgram(ui.InitialTextInputModel(ui.TextInputOptions{
				Label:       "Task name:",
				Placeholder: "Enter the name of the task",
				CharLimit:   50,
				Required:    true,
			}, &output))
			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}

			name = output.Value
		}

		if description == "" {
			p := tea.NewProgram(ui.InitialTextInputModel(ui.TextInputOptions{
				Label:       "Task description:",
				Placeholder: "Enter the description of the task. Press Enter to skip",
				CharLimit:   256,
				Required:    false,
			}, &output))
			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}

			if output.Value != "" {
				description = output.Value
			} else {
				description = "No description"
			}
		}

		if category == "" {
			p := tea.NewProgram(ui.InitialTextInputModel(ui.TextInputOptions{
				Label:       "Task category:",
				Placeholder: "Enter the category of the task. Press Enter to skip",
				CharLimit:   30,
				Required:    false,
			}, &output))
			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}

			if output.Value != "" {
				category = output.Value
			} else {
				category = "general"
			}
		}

		TodoList.AddTask(name, description, category)
		fmt.Println(config.SuccessStyle.Render("Task added successfully"))
	},
}

func init() {
	addCmd.Flags().StringP("name", "n", "", "Name of the task")
	addCmd.Flags().StringP("description", "d", "", "Description of the task")
	addCmd.Flags().StringP("category", "c", "", "Name of the category to which the task belongs")

}
