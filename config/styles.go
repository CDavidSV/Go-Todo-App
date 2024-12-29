package config

import "github.com/charmbracelet/lipgloss"

var (
	SuccessColor string = "#00ff2a"
	ErrorColor   string = "#ff0000"
	InfoColor    string = "#00a2ff"
	WarningColor string = "#ff8c00"

	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(SuccessColor)).
			Italic(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(ErrorColor))

	InfoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(InfoColor))

	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(WarningColor))

	TableBorderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("99"))

	TableHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("99")).
				Bold(true).
				Align(lipgloss.Center)

	TableEvenRowStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#cccccc"))

	TableOddRowStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#999999"))

	Logo = `
 ______    _____       ____       _____
/\__  _\  /\  __ \    /\  _ \    /\  __ \
\/_/\ \/  \ \ \/\ \   \ \ \/\ \  \ \ \/\ \
   \ \ \   \ \ \ \ \   \ \ \ \ \  \ \ \ \ \
    \ \ \   \ \ \_\ \   \ \ \_\ \  \ \ \_\ \
     \ \_\   \ \_____\   \ \____/   \ \_____\
      \/_/    \/_____/    \/___/     \/_____/`
)
