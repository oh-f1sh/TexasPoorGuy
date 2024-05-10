package client

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type GameModel struct {
	viewport viewport.Model
	messages []string
}

func InitialGameModel() GameModel {
	vp := viewport.New(30, 20)
	message := []string{}
	return GameModel{
		viewport: vp,
		messages: message,
	}
}

func (m GameModel) Init() tea.Cmd {
	return nil
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.viewport.SetContent(strings.Join(m.messages, "\n"))
	m.viewport.GotoBottom()
	return m, nil
}

func (m GameModel) View() string {
	return m.viewport.View()
}
