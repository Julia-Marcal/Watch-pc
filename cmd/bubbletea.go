package cmd

import (
	"fmt"
	"strings"

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
	var cmd []tea.Cmd = make([]tea.Cmd, len(p.Inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if p.focused == len(p.Inputs)-1 {
				return p, tea.Quit
			}
			p.nextInput()
		case tea.KeyCtrlC, tea.KeyEsc:
			return p, tea.Quit
		case tea.KeyShiftTab, tea.KeyCtrlP:
			p.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
			p.nextInput()
		}
		for i := range p.Inputs {
			p.Inputs[i].Blur()
		}
		p.Inputs[p.focused].Focus()

	case errMsg:
		p.err = msg
		return p, nil
	}

	for i := range p.Inputs {
		p.Inputs[i], cmd[i] = p.Inputs[i].Update(msg)
	}
	return p, tea.Batch(cmd...)
}

func (p CmdModel) View() string {
	s := strings.Builder{}
	s.WriteString("Type your command")
	return fmt.Sprintf(
		`
		%s
		
		`,
		p.Inputs[Command].View(),
	) + "\n"
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
