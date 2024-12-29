package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

type TextInputModel struct {
	textInput textinput.Model
	label     string
	err       error
	output    *TextInputValue
	req       bool
}

type TextInputOptions struct {
	Label       string
	Placeholder string
	CharLimit   int
	Required    bool
}

type TextInputValue struct {
	Value string
}

type errMsg error

var labelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#fffff")).Background(lipgloss.Color("99")).Padding(0, 1)

func InitialTextInputModel(options TextInputOptions, output *TextInputValue) TextInputModel {
	ti := textinput.New()

	ti.Placeholder = options.Placeholder
	ti.CharLimit = options.CharLimit
	ti.Focus()

	if options.CharLimit < 1 {
		ti.CharLimit = 256
	}

	output.Value = ""

	return TextInputModel{
		textInput: ti,
		err:       nil,
		label:     options.Label,
		output:    output,
		req:       options.Required,
	}
}

func (m TextInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TextInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.req && m.textInput.Value() == "" {
				return m, nil
			}

			m.output.Value = m.textInput.Value()

			return m, tea.Quit
		case tea.KeyCtrlC, tea.KeyEsc:
			os.Exit(0)
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m TextInputModel) View() string {
	return fmt.Sprintf(
		"%s\n%s\n\n",
		labelStyle.Render(m.label),
		m.textInput.View(),
	)
}
