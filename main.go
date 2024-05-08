package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/oh-f1sh/TexasPoorGuy/client"
)

func main() {
	p := tea.NewProgram(client.InitialPoorGuyClient())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
