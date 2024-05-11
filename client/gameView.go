package client

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type gameUpdate struct {
	msg string
}

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
	return waitForGameMsg()
}

func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var vpCmd tea.Cmd
	switch msgTyp := msg.(type) {
	case tea.KeyMsg:
		switch msgTyp.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case gameUpdate:
		m.messages = append(m.messages, msg.(gameUpdate).msg)
		m.viewport.SetContent(strings.Join(m.messages, "\n"))
		m.viewport.GotoBottom()
		m.viewport, vpCmd = m.viewport.Update(msg)
	}
	return m, tea.Batch(vpCmd, waitForGameMsg())
}

func (m GameModel) View() string {
	return m.viewport.View()
}

func waitForGameMsg() tea.Cmd {
	return func() tea.Msg {
		return gameUpdate{
			msg: "∙ " + <-GAME_MSG_CHAN,
		}
	}
}
