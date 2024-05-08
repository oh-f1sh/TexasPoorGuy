package client

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type ScoreBoardModel struct {
	viewport viewport.Model
	scores   []string
}

func InitialScoreBoardModel() ScoreBoardModel {
	vp := viewport.New(30, 5)
	scores := []string{
		"MTT: Score 100",
		"MTT: Score 100",
		"MTT: Score 100",
		"MTT: Score 100",
	}
	vp.SetContent(strings.Join(scores, "\n"))
	return ScoreBoardModel{
		viewport: vp,
		scores:   scores,
	}
}

func (m ScoreBoardModel) Init() tea.Cmd {
	m.viewport.SetContent(strings.Join(m.scores, "\n"))
	return nil
}

func (m ScoreBoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.viewport.SetContent(strings.Join(m.scores, "\n"))
	return m, nil
}

func (m ScoreBoardModel) View() string {
	return m.viewport.View()
}
