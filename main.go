package main

import (
	"fmt"
	"os"

	cmd "github.com/Julia-Marcal/watch-pc-cmd/cmd"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		cmd.CmdWatchPc(),
	)

	m, err := p.Run()

	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(cmd.ValorantModel); ok {
		commandValue := m.Inputs[cmd.Command].Value()
	}

}
