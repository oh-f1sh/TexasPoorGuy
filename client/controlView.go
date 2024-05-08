package client

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var Choices = []string{
	"call", "raise", "allin", "fold", "check",
}

type ControlModel struct {
	cursor int
	choice string
}

func InitialControlModel() ControlModel {
	return ControlModel{
		cursor: 0,
		choice: "",
	}
}

func (m ControlModel) Init() tea.Cmd {
	return nil
}

func (m ControlModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.choice = Choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(Choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(Choices) - 1
			}
		}
	}

	return m, nil
}

func (m ControlModel) View() string {
	s := strings.Builder{}
	s.WriteString("Use arrow keys to navigate between options.\n\n")

	for i := 0; i < len(Choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(Choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press enter to confirm)\n")

	return s.String()
}
