package main

import (
	_ "github.com/oh-f1sh/TexasPoorGuy/conf"

	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/oh-f1sh/TexasPoorGuy/client"
)

func main() {
	p := tea.NewProgram(client.InitialLoginModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
