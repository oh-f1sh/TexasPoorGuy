package client

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(0, 2)

type Room struct {
	title string
	desc  string
}

func (r Room) Title() string       { return r.title }
func (r Room) Description() string { return r.desc }
func (r Room) FilterValue() string { return r.title }

type RoomModel struct {
	list list.Model
}

func InitialRoomModel() RoomModel {
	items := []list.Item{
		Room{title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		Room{title: "Nutella", desc: "It's good on toast"},
		Room{title: "Bitter melon", desc: "It cools you down"},
		Room{title: "Nice socks", desc: "And by that I mean socks without holes"},
		Room{title: "Eight hours of sleep", desc: "I had this once"},
		Room{title: "Cats", desc: "Usually"},
		Room{title: "Plantasia, the album", desc: "My plants love it too"},
		Room{title: "Pour over coffee", desc: "It takes forever to make though"},
		Room{title: "VR", desc: "Virtual reality...what is there to say?"},
		Room{title: "Noguchi Lamps", desc: "Such pleasing organic forms"},
		Room{title: "Linux", desc: "Pretty much the best OS"},
		Room{title: "Business school", desc: "Just kidding"},
		Room{title: "Pottery", desc: "Wet clay is a great feeling"},
		Room{title: "Shampoo", desc: "Nothing like clean hair"},
		Room{title: "Table tennis", desc: "It’s surprisingly exhausting"},
		Room{title: "Milk crates", desc: "Great for packing in your extra stuff"},
		Room{title: "Afternoon tea", desc: "Especially the tea sandwich part"},
		Room{title: "Stickers", desc: "The thicker the vinyl the better"},
		Room{title: "20° Weather", desc: "Celsius, not Fahrenheit"},
		Room{title: "Warm light", desc: "Like around 2700 Kelvin"},
		Room{title: "The vernal equinox", desc: "The autumnal equinox is pretty good too"},
		Room{title: "Gaffer’s tape", desc: "Basically sticky fabric"},
		Room{title: "Terrycloth", desc: "In other words, towel fabric"},
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Texas Poor Guy - Room List"
	m := RoomModel{list: l}

	return m
}

func (m RoomModel) Init() tea.Cmd {
	return nil
}

func (m RoomModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			fmt.Println("chose", m.list.SelectedItem())
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m RoomModel) View() string {
	return docStyle.Render(m.list.View())
}
