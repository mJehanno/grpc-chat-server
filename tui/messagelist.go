package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type messageListModel struct {
	messages []Message
}

func (m messageListModel) Init() tea.Cmd {
	return nil
}

func (m messageListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case newIncomingMessage:
		m.messages = append(m.messages, Message{author: msg.Author, content: msg.Content})
		return m, nil
	}
	return nil, nil
}

func (m messageListModel) View() string {
	s := ""
	for _, msg := range m.messages {
		s += fmt.Sprintf("%s : %s \n", msg.author, msg.content)
	}
	return s
}
