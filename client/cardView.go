package client

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type communityCardUpdate struct {
	card  []string
	color []string
	bg    []string
}

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
	return waitForCommunityCard()
}

func (m CardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgTyp := msg.(type) {
	case tea.KeyMsg:
		switch msgTyp.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case communityCardUpdate:
		cards := msg.(communityCardUpdate)
		for i := range m.card {
			m.viewport[i].SetContent(cards.card[i])
			m.viewport[i].Style = lipgloss.NewStyle().
				Background(lipgloss.Color(cards.bg[i])).
				Foreground(lipgloss.Color(cards.color[i])).
				Align(lipgloss.Center).
				Margin(1).
				Bold(true)
		}
	}
	return m, waitForCommunityCard()
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

func waitForCommunityCard() tea.Cmd {
	return func() tea.Msg {
		return <-COMMUNITY_CARD_CHAN
	}
}

type handCardUpdate struct {
	card  []string
	color []string
	bg    []string
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
	return waitForHandCard()
}

func (m HandCardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msgTyp := msg.(type) {
	case tea.KeyMsg:
		switch msgTyp.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case handCardUpdate:
		cards := msg.(handCardUpdate)
		m.viewport[card1].SetContent(cards.card[card1])
		m.viewport[card1].Style = lipgloss.NewStyle().
			Background(lipgloss.Color(cards.bg[card1])).
			Foreground(lipgloss.Color(cards.color[card1])).
			Align(lipgloss.Center).
			Margin(1).
			Bold(true)
		m.viewport[card2].SetContent(cards.card[card2])
		m.viewport[card2].Style = lipgloss.NewStyle().
			Background(lipgloss.Color(cards.bg[card2])).
			Foreground(lipgloss.Color(cards.color[card2])).
			Align(lipgloss.Center).
			Margin(1).
			Bold(true)
	}

	return m, waitForHandCard()
}

func (m HandCardModel) View() string {
	return lipgloss.JoinHorizontal(lipgloss.Top,
		m.viewport[card1].View(),
		m.viewport[card2].View(),
	)
}

func waitForHandCard() tea.Cmd {
	return func() tea.Msg {
		return <-HAND_CARD_CHAN
	}
}
