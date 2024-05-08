package client

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	UnfocusedStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("243"))
	FocusedStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("69"))
	OtherStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(hotPink))
)

const (
	focusChat = iota
	focusControl
)

type PoorGuyClient struct {
	CardView       CardModel
	HandCardView   HandCardModel
	GameView       GameModel
	ControlView    ControlModel
	ScoreBoardView ScoreBoardModel
	ChatView       ChatModel
	Focus          int
}

func InitialPoorGuyClient() PoorGuyClient {
	return PoorGuyClient{
		CardView:       InitialCardModel(),
		HandCardView:   InitialHandCardModel(),
		GameView:       InitialGameModel(),
		ControlView:    InitialControlModel(),
		ScoreBoardView: InitialScoreBoardModel(),
		ChatView:       InitialChatModel(),
		Focus:          focusControl,
	}
}

func (m PoorGuyClient) Init() tea.Cmd {
	return nil
}

func (m PoorGuyClient) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab:
			if m.Focus == focusChat {
				m.Focus = focusControl
			} else {
				m.Focus = focusChat
			}
		}
		switch m.Focus {
		case focusChat:
			m.ChatView.Update(msg)
		case focusControl:
			m.ControlView.Update(msg)
		}
	}

	return m, nil
}

func (m PoorGuyClient) View() string {
	if m.Focus == focusChat {
		return lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.JoinVertical(lipgloss.Left,
				lipgloss.JoinHorizontal(lipgloss.Top,
					OtherStyle.Render(m.CardView.View()),
					OtherStyle.Render(m.HandCardView.View()),
				),
				OtherStyle.Render(m.GameView.View()),
				UnfocusedStyle.Render(m.ControlView.View()),
			),
			lipgloss.JoinVertical(lipgloss.Left,
				OtherStyle.Render(m.ScoreBoardView.View()),
				FocusedStyle.Render(m.ChatView.View()),
			),
		)
	} else {
		return lipgloss.JoinHorizontal(lipgloss.Top,
			lipgloss.JoinVertical(lipgloss.Left,
				lipgloss.JoinHorizontal(lipgloss.Top,
					OtherStyle.Render(m.CardView.View()),
					OtherStyle.Render(m.HandCardView.View()),
				),
				OtherStyle.Render(m.GameView.View()),
				FocusedStyle.Render(m.ControlView.View()),
			),
			lipgloss.JoinVertical(lipgloss.Left,
				OtherStyle.Render(m.ScoreBoardView.View()),
				UnfocusedStyle.Render(m.ChatView.View()),
			),
		)
	}
}
