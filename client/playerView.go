package client

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	greenHighlight = "150"
)

var (
	activePlayerStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(greenHighlight)).
				Foreground(lipgloss.Color(darkgrey)).
				Align(lipgloss.Center).
				Margin(1).
				Bold(true)

	inactivePlayerStyle = lipgloss.NewStyle().
				Background(lipgloss.Color(darkgrey)).
				Foreground(lipgloss.Color(white)).
				Align(lipgloss.Center).
				Margin(1).
				Bold(false)
)

type PlayerModel struct {
	viewport []viewport.Model
	players  []string
	state    int
}

func InitialPlayerModel() PlayerModel {
	vp := []viewport.Model{
		viewport.New(10, 2),
		viewport.New(10, 2),
		viewport.New(10, 2),
		viewport.New(10, 2),
		viewport.New(10, 2),
		viewport.New(10, 2),
	}
	players := []string{
		"MTT\nBig Blind",
		"szy\nSmall Blind",
		"kk",
		"Bovia",
		"lhr",
		"xu",
	}
	for i, p := range players {
		vp[i].SetContent(p)
	}
	return PlayerModel{
		viewport: vp,
		players:  players,
		state:    1,
	}
}

func (m PlayerModel) Init() tea.Cmd {
	return nil
}

func (m PlayerModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m PlayerModel) View() string {
	if len(m.players) <= 5 {
		line := make([]string, 0)
		for i := 0; i < len(m.players); i++ {
			if i == m.state {
				m.viewport[i].Style = activePlayerStyle
			} else {
				m.viewport[i].Style = inactivePlayerStyle
			}
			line = append(line, m.viewport[i].View())
		}
		return lipgloss.JoinHorizontal(lipgloss.Top, line...)
	} else {
		line1 := make([]string, 0)
		line2 := make([]string, 0)

		for i := 0; i < 5; i++ {
			if i == m.state {
				m.viewport[i].Style = activePlayerStyle
			} else {
				m.viewport[i].Style = inactivePlayerStyle
			}
			line1 = append(line1, m.viewport[i].View())
		}
		for i := 5; i < len(m.players); i++ {
			if i == m.state {
				m.viewport[i].Style = activePlayerStyle
			} else {
				m.viewport[i].Style = inactivePlayerStyle
			}
			line2 = append(line2, m.viewport[i].View())
		}
		return lipgloss.JoinVertical(lipgloss.Left,
			lipgloss.JoinHorizontal(lipgloss.Top, line1...),
			lipgloss.JoinHorizontal(lipgloss.Top, line2...),
		)
	}
}
