package client

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(0, 2)
var refreshTimeStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(darkgrey))
var additonalHelpStyle = lipgloss.NewStyle().
	Foreground(hotPink).
	MarginLeft(4)

var ROOM_LIST = make([]Room, 0)

type Room struct {
	id          int
	ownerID     int
	ownerName   string
	gameType    string
	playerCount int
}

func (r Room) Title() string { return fmt.Sprintf("%s's Game Room", r.ownerName) }
func (r Room) Description() string {
	return fmt.Sprintf("game type: %s, now %d players", r.gameType, r.playerCount)
}
func (r Room) FilterValue() string { return r.Title() + r.Description() }

type RoomModel struct {
	list list.Model
}

func InitialRoomModel() RoomModel {
	items := []list.Item{}
	ListRoom()
	<-LIST_ROOM_SIGNAL
	for _, v := range ROOM_LIST {
		items = append(items, v)
	}
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Texas Poor Guy - Room List"
	l.SetSize(60, 40)
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
		case tea.KeyEsc, tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			if len(ROOM_LIST) > 0 {
				JoinRoom(ROOM_LIST[m.list.Index()].id)
				if <-JOIN_ROOM_SIGNAL == 1 {
					cli := InitialPoorGuyClient()
					return cli, cli.Init()
				}
			}
		}
		switch msg.String() {
		case "n":
			CreateRoom()
			if <-CREATE_ROOM_SIGNAL == 1 {
				cli := InitialPoorGuyClient()
				return cli, cli.Init()
			}
		case "r":
			ListRoom()
			<-LIST_ROOM_SIGNAL
			items := []list.Item{}
			for _, v := range ROOM_LIST {
				items = append(items, v)
			}
			m.list.SetItems(items)
		case "p":
			GetSubsidy()
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m RoomModel) View() string {
	return docStyle.Render(m.list.View()) +
		"\n\n" + additonalHelpStyle.Render("Enter	join room	|	n	new room\nr		refresh list	|	p	get subsidy"+
		"\n\n"+refreshTimeStyle.Render("last refreshed at: "+time.Now().Format("15:04:05")))
}
