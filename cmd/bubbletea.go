package cmd

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

const (
	Command = iota
)

type CmdModel struct {
	Inputs  []textinput.Model
	focused int
	err     error
}

func (p CmdModel) Init() tea.Cmd {
	return nil
}

func CmdWatchPc() CmdModel {
	var Inputs []textinput.Model = make([]textinput.Model, 1)

	Inputs[Command] = textinput.New()
	Inputs[Command].Placeholder = "Your command"
	Inputs[Command].Focus()
	Inputs[Command].CharLimit = 16
	Inputs[Command].Width = 30

	return CmdModel{
		Inputs:  Inputs,
		focused: 0,
		err:     nil,
	}

}

func (p CmdModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (p CmdModel) View() string {
	return "nil"
}

func (p *CmdModel) nextInput() {
	p.focused = (p.focused + 1) % len(p.Inputs)
}

func (p *CmdModel) prevInput() {
	p.focused--
	if p.focused < 0 {
		p.focused = len(p.Inputs) - 1
	}
}
