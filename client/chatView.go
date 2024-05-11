package client

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/oh-f1sh/TexasPoorGuy/common"
)

type chatUpdate struct {
	content string
}

type ChatModel struct {
	viewport         viewport.Model
	textarea         textarea.Model
	senderStyle      lipgloss.Style
	otherSenderStyle lipgloss.Style
}

func InitialChatModel() ChatModel {
	ta := textarea.New()
	ta.Placeholder = "Send a message..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = 280

	ta.SetWidth(30)
	ta.SetHeight(3)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	vp := viewport.New(30, 19)
	vp.SetContent(`Welcome to the chat room!
Type a message and press Enter to send.`)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return ChatModel{
		textarea:         ta,
		viewport:         vp,
		senderStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("5")),
		otherSenderStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("11")),
	}
}

func (m ChatModel) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		waitForIncomingMessage(),
	)
}

func (m ChatModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)

	switch msgTyp := msg.(type) {
	case tea.KeyMsg:
		switch msgTyp.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			SendMsg(common.ROOMID, m.textarea.Value())
			m.textarea.Reset()
		}
	case chatUpdate:
		m.viewport.SetContent(msg.(chatUpdate).content)
		m.viewport.GotoBottom()
		m.viewport, vpCmd = m.viewport.Update(msg)
		return m, tea.Batch(tiCmd, vpCmd, waitForIncomingMessage())
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m ChatModel) View() string {
	return fmt.Sprintf(
		"%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
	) + "\n\n"
}

func waitForIncomingMessage() tea.Cmd {
	return func() tea.Msg {
		<-SEND_ROOM_MSG_SIGNAL
		return chatUpdate{
			content: common.ROOM_CHAT,
		}
	}
}
