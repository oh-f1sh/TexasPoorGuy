package client

import (
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/oh-f1sh/TexasPoorGuy/common"
)

var Choices = []string{
	"call", "raise", "allin", "fold", "check",
}

type ControlModel struct {
	textarea   textarea.Model
	cursor     int
	choice     string
	raiseValue int
}

func InitialControlModel() ControlModel {
	ta := textarea.New()
	ta.Placeholder = "e.g. 50"

	ta.Prompt = ""
	ta.CharLimit = 100
	ta.SetWidth(50)
	ta.SetHeight(1)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)

	return ControlModel{
		textarea:   ta,
		cursor:     0,
		choice:     "",
		raiseValue: 0,
	}
}

func (m ControlModel) Init() tea.Cmd {
	return textarea.Blink
}

func (m ControlModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.choice == "raise" {
				raiseValue, _ := strconv.Atoi(m.textarea.Value())
				UserAction(m.choice, raiseValue)
				m.choice = ""
				return m, nil
			}
			m.choice = Choices[m.cursor]
		case tea.KeyDown:
			m.cursor++
			if m.cursor >= len(Choices) {
				m.cursor = 0
			}
		case tea.KeyUp:
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(Choices) - 1
			}
		}
		switch msg.String() {
		case "e":
			// exit room
			return InitialRoomModel(), nil
		case "s":
			// owner can start game
			if common.USERID == common.ROOMOWNERID {
				StartGame()
			}
		}
	}

	switch m.choice {
	case "raise":
		m.textarea.Focus()
		m.textarea.Update(msg)
		m.textarea.SetCursor(100)
		return m, nil
	case "call":
		UserAction(m.choice, 0)
		m.choice = ""
	case "allin":
		UserAction(m.choice, 0)
		m.choice = ""
	case "fold":
		UserAction(m.choice, 0)
		m.choice = ""
	case "check":
		UserAction(m.choice, 0)
		m.choice = ""
	}

	return m, nil
}

func (m ControlModel) View() string {
	if m.choice == "raise" {
		s := strings.Builder{}
		s.WriteString("You chose to raise the bet, enter the amount.\n\n\n\n\n")
		s.WriteString(m.textarea.View())
		s.WriteString("\n\n\n(press enter to confirm)\n")
		return s.String()
	} else {
		s := strings.Builder{}
		s.WriteString("Use arrow keys to navigate between options.       \n\n")

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
}
