package client

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var waitingPlayerStyle = lipgloss.NewStyle().Background(lipgloss.Color(darkgrey))
var waitingPlayerNameStyle = lipgloss.NewStyle().Foreground(hotPink).Background(lipgloss.Color(darkgrey))
var playingPlayerStyle = lipgloss.NewStyle().Background(lipgloss.Color(lightgrey))
var playingPlayerNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(lightpink)).Background(lipgloss.Color(lightgrey))

type scoreboardUpdate struct {
	scores []string
}

type ScoreBoardModel struct {
	viewport viewport.Model
	scores   []string
}

func InitialScoreBoardModel() ScoreBoardModel {
	vp := viewport.New(30, 12)
	scores := []string{}
	vp.SetContent(strings.Join(scores, "\n"))
	return ScoreBoardModel{
		viewport: vp,
		scores:   scores,
	}
}

func (m ScoreBoardModel) Init() tea.Cmd {
	return waitForScoreboard()
}

func (m ScoreBoardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgTyp := msg.(type) {
	case tea.KeyMsg:
		switch msgTyp.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case scoreboardUpdate:
		scores := msg.(scoreboardUpdate).scores
		m.viewport.SetContent(strings.Join(scores, "\n"))
	}
	return m, waitForScoreboard()
}

func (m ScoreBoardModel) View() string {
	return m.viewport.View()
}

func waitForScoreboard() tea.Cmd {
	return func() tea.Msg {
		return <-SCOREBOARD_CHAN
	}
}
