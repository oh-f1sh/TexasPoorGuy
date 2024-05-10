package client

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CardModel struct {
	viewport   []viewport.Model
	card       []string
	color      []string
	background []string
}

var (
	suitMap = map[int]string{
		0: "♠",
		1: "♥",
		2: "♣",
		3: "♦",
	}
	suitColor = map[int]string{
		0: black,
		1: red,
		2: black,
		3: red,
	}
)

const (
	card1 = iota
	card2
	card3
	card4
	card5

	darkgrey = "8"
	black    = "16"
	red      = "1"
	white    = "255"
)

func InitialCardModel() CardModel {
	vp := []viewport.Model{
		viewport.New(4, 4),
		viewport.New(4, 4),
		viewport.New(4, 4),
		viewport.New(4, 4),
		viewport.New(4, 4),
	}
	bg := []string{darkgrey, darkgrey, darkgrey, darkgrey, darkgrey}
	clr := []string{black, black, black, black, black}
	card := []string{"", "", "", "", ""}
	for i := range card {
		vp[i].SetContent(card[i])
		vp[i].Style = lipgloss.NewStyle().
			Background(lipgloss.Color(bg[i])).
			Foreground(lipgloss.Color(clr[i])).
			Align(lipgloss.Center).
			Margin(1).
			Bold(true)
	}

	return CardModel{
		viewport:   vp,
		card:       card,
		color:      clr,
		background: bg,
	}
}

func (m CardModel) Init() tea.Cmd {
	return nil
}

func (m CardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	for i := range m.card {
		m.viewport[i].SetContent(m.card[i])
		m.viewport[i].Style = lipgloss.NewStyle().
			Background(lipgloss.Color(m.background[i])).
			Foreground(lipgloss.Color(m.color[i])).
			Align(lipgloss.Center).
			Margin(1).
			Bold(true)
	}
	return m, nil
}

func (m CardModel) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		m.viewport[card1].View(),
		m.viewport[card2].View(),
		m.viewport[card3].View(),
		m.viewport[card4].View(),
		m.viewport[card5].View(),
	)
}

type HandCardModel struct {
	viewport   []viewport.Model
	card       []string
	color      []string
	background []string
}

func InitialHandCardModel() HandCardModel {
	vp := []viewport.Model{
		viewport.New(4, 4),
		viewport.New(4, 4),
	}
	bg := []string{darkgrey, darkgrey}
	clr := []string{black, black}
	card := []string{"", ""}
	vp[card1].SetContent(card[card1])
	vp[card2].SetContent(card[card2])

	vp[card1].Style = lipgloss.NewStyle().
		Background(lipgloss.Color(bg[card1])).
		Foreground(lipgloss.Color(clr[card1])).
		Align(lipgloss.Center).
		Margin(1).
		Bold(true)
	vp[card2].Style = lipgloss.NewStyle().
		Background(lipgloss.Color(bg[card2])).
		Foreground(lipgloss.Color(clr[card2])).
		Align(lipgloss.Center).
		Margin(1).
		Bold(true)

	return HandCardModel{
		viewport:   vp,
		card:       card,
		color:      clr,
		background: bg,
	}
}

func (m HandCardModel) Init() tea.Cmd {
	return nil
}

func (m HandCardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.viewport[card1].SetContent(m.card[card1])
	m.viewport[card1].Style = lipgloss.NewStyle().
		Background(lipgloss.Color(m.background[card1])).
		Foreground(lipgloss.Color(m.color[card1])).
		Align(lipgloss.Center).
		Margin(1).
		Bold(true)
	m.viewport[card2].SetContent(m.card[card2])
	m.viewport[card2].Style = lipgloss.NewStyle().
		Background(lipgloss.Color(m.background[card2])).
		Foreground(lipgloss.Color(m.color[card2])).
		Align(lipgloss.Center).
		Margin(1).
		Bold(true)
	return m, nil
}

func (m HandCardModel) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		m.viewport[card1].View(),
		m.viewport[card2].View(),
	)
}
