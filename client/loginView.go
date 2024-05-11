package client

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/oh-f1sh/TexasPoorGuy/common"
)

type LoginModel struct {
	inputs  []textinput.Model
	focused int
}

// Init implements tea.Model.
func (l LoginModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements tea.Model.
func (l LoginModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(l.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if l.focused == len(l.inputs)-1 {
				Login(l.inputs[name].Value(), l.inputs[pwd].Value(), l.inputs[address].Value())
				isLogin := <-LOGIN_SIGNAL
				if isLogin == 1 {
					return InitialRoomModel(), tea.Batch(cmds...)
				}
			}
			l.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return l, tea.Quit
		case tea.KeyShiftTab:
			l.prevInput()
		case tea.KeyTab:
			l.nextInput()
		}
		for i := range l.inputs {
			l.inputs[i].Blur()
		}
		l.inputs[l.focused].Focus()
	}

	for i := range l.inputs {
		l.inputs[i], cmds[i] = l.inputs[i].Update(msg)
	}
	return l, tea.Batch(cmds...)
}

// View implements tea.Model.
func (l LoginModel) View() string {
	return fmt.Sprintf(
		` Welcome to Texas Poor Guy

 %s
 %s

 %s
 %s


 %s
 %s

 %s
`,
		inputStyle.Width(30).Render("User Name"),
		l.inputs[name].View(),
		inputStyle.Width(30).Render("Password"),
		l.inputs[pwd].View(),
		inputStyle.Width(20).Render("Server Address"),
		l.inputs[address].View(),
		continueStyle.Render("Press Enter to start Texas Poor Guy!    =========>"),
	) + "\n"
}

const (
	name = iota
	pwd
	address
)

var (
	hotPink       = lipgloss.Color("#FF06B7")
	darkGray      = lipgloss.Color("#767676")
	inputStyle    = lipgloss.NewStyle().Foreground(hotPink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)

func InitialLoginModel() LoginModel {
	var inputs []textinput.Model = make([]textinput.Model, 3)
	inputs[name] = textinput.New()
	inputs[name].Placeholder = "Your user name"
	inputs[name].Focus()
	inputs[name].CharLimit = 20
	inputs[name].Width = 30
	inputs[name].Prompt = ""
	if len(common.DEFAULTUSERNAME) > 0 {
		inputs[name].SetValue(common.DEFAULTUSERNAME)
	}

	inputs[pwd] = textinput.New()
	inputs[pwd].Placeholder = "Your password"
	inputs[pwd].CharLimit = 20
	inputs[pwd].Width = 30
	inputs[pwd].Prompt = ""
	inputs[pwd].EchoMode = textinput.EchoPassword
	inputs[pwd].EchoCharacter = '•'
	if len(common.DEFAULTPWD) > 0 {
		inputs[pwd].SetValue(common.DEFAULTPWD)
	}

	inputs[address] = textinput.New()
	inputs[address].Placeholder = "xxx.xxx.xxx.xxx "
	inputs[address].CharLimit = 15
	inputs[address].Width = 20
	inputs[address].Prompt = ""
	if len(common.SERVER_ADDR) > 0 {
		inputs[address].SetValue(common.SERVER_ADDR)
	}

	return LoginModel{
		inputs:  inputs,
		focused: 0,
	}
}

// nextInput focuses the next input field
func (l *LoginModel) nextInput() {
	l.focused = (l.focused + 1) % len(l.inputs)
}

// prevInput focuses the previous input field
func (l *LoginModel) prevInput() {
	l.focused--
	// Wrap around
	if l.focused < 0 {
		l.focused = len(l.inputs) - 1
	}
}
