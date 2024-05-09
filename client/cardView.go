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
	bg := []string{"255", "255", "255", "255", "255"}
	clr := []string{"16", "1", "1", "16", "1"}
	card := []string{"2\nX", "J\nX", "Q\nX", "K\nX", "10\nX"}
	vp[card1].SetContent(card[card1])
	vp[card2].SetContent(card[card2])
	vp[card3].SetContent(card[card3])
	vp[card4].SetContent(card[card4])
	vp[card5].SetContent(card[card5])
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
	vp[card3].Style = lipgloss.NewStyle().
		Background(lipgloss.Color(bg[card3])).
		Foreground(lipgloss.Color(clr[card3])).
		Align(lipgloss.Center).
		Margin(1).
		Bold(true)
	vp[card4].Style = lipgloss.NewStyle().
		Background(lipgloss.Color(bg[card4])).
		Foreground(lipgloss.Color(clr[card4])).
		Align(lipgloss.Center).
		Margin(1).
		Bold(true)
	vp[card5].Style = lipgloss.NewStyle().
		Background(lipgloss.Color(bg[card5])).
		Foreground(lipgloss.Color(clr[card5])).
		Align(lipgloss.Center).
		Margin(1).
		Bold(true)

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
	bg := []string{white, white}
	clr := []string{red, black}
	card := []string{"A\nX", "K\nX"}
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
	return m, nil
}

func (m HandCardModel) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		m.viewport[card1].View(),
		m.viewport[card2].View(),
	)
}
