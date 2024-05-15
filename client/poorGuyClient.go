package client

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var POOR_GUY_CLIENT *PoorGuyClient

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
	cli := PoorGuyClient{
		CardView:       InitialCardModel(),
		HandCardView:   InitialHandCardModel(),
		GameView:       InitialGameModel(),
		ControlView:    InitialControlModel(),
		ScoreBoardView: InitialScoreBoardModel(),
		ChatView:       InitialChatModel(),
		Focus:          focusChat,
	}
	POOR_GUY_CLIENT = &cli
	return cli
}

func (m PoorGuyClient) Init() tea.Cmd {
	return tea.Batch(
		m.CardView.Init(),
		m.HandCardView.Init(),
		m.GameView.Init(),
		m.ControlView.Init(),
		m.ScoreBoardView.Init(),
		m.ChatView.Init(),
	)
}

func (m PoorGuyClient) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := make([]tea.Cmd, 0)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

		// switch window focus
		case tea.KeyTab:
			if m.Focus == focusChat {
				m.Focus = focusControl
			} else {
				m.Focus = focusChat
			}
		}
		switch m.Focus {
		case focusChat:
			c, cmd := m.ChatView.Update(msg)
			m.ChatView = c.(ChatModel)
			cmds = append(cmds, cmd)
		case focusControl:
			if msg.String() == "e" {
				QuitRoom()
				return InitialRoomModel(), nil
			}
			c, cmd := m.ControlView.Update(msg)
			m.ControlView = c.(ControlModel)
			cmds = append(cmds, cmd)
		}
	case chatUpdate:
		c, cmd := m.ChatView.Update(msg)
		m.ChatView = c.(ChatModel)
		cmds = append(cmds, cmd)
	case gameUpdate:
		c, cmd := m.GameView.Update(msg)
		m.GameView = c.(GameModel)
		cmds = append(cmds, cmd)
	case handCardUpdate:
		c, cmd := m.HandCardView.Update(msg)
		m.HandCardView = c.(HandCardModel)
		cmds = append(cmds, cmd)
	case communityCardUpdate:
		c, cmd := m.CardView.Update(msg)
		m.CardView = c.(CardModel)
		cmds = append(cmds, cmd)
	case scoreboardUpdate:
		c, cmd := m.ScoreBoardView.Update(msg)
		m.ScoreBoardView = c.(ScoreBoardModel)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
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
