package main

import (
	"fmt"
	"os"

	cmd "github.com/Julia-Marcal/Watch-pc/cmd"
	pc "github.com/Julia-Marcal/Watch-pc/internal"

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

	if m, ok := m.(cmd.CmdModel); ok {
		commandValue := m.Inputs[cmd.Command].Value()
	}

}
