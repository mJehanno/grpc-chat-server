package tui

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mjehanno/grpc-chat/service/chat"
	"google.golang.org/grpc"
)

type newIncomingMessage struct {
	Author  string
	Content string
}

type clientCreatedMessage struct {
	client chat.ChatServiceClient
}

type Message struct {
	author  string
	content string
}

type model struct {
	inputs      []textinput.Model
	messageList messageListModel
	focusIndex  int
	messageChan chan Message
	client      chat.ChatServiceClient
}

func InitialModel() model {
	ti := textinput.New()
	ti.Placeholder = "message"

	user := textinput.New()
	user.Placeholder = "username"
	user.Focus()

	return model{
		inputs:      []textinput.Model{user, ti},
		messageChan: make(chan Message),
		focusIndex:  0,
	}
}

func (m model) Init() tea.Cmd {
	tea.Println("init")
	return tea.Batch(startGrpcClient(m.messageChan), readFromGrpc(m.messageChan))
}

func (m *model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func readFromGrpc(messageChan chan Message) tea.Cmd {
	fmt.Println("-------------------")
	fmt.Println("readFromGrpc")
	go func() tea.Cmd {
		for message := range messageChan {
			fmt.Println("received a message")
			return func() tea.Msg {
				return newIncomingMessage{Author: message.author, Content: message.content}
			}
		}
		fmt.Println("toto")
		return nil
	}()
	return nil
}

func startGrpcClient(messageChan chan Message) tea.Cmd {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client cannot connect to server : %s", err.Error())
	}

	client := chat.NewChatServiceClient(conn)

	ctx := context.Background()

	messageStream, err := client.ReceiveMsg(ctx, &chat.Empty{})
	if err != nil {
		log.Fatal("failed to connect to grpc server")
	}

	go func() {
		for {
			message, err := messageStream.Recv()
			fmt.Println("message received before transmitting to channel")
			if err == io.EOF {
				close(messageChan)
				return
			}
			messageChan <- Message{author: message.Author, content: message.Content}
		}
	}()

	return func() tea.Msg {
		return clientCreatedMessage{
			client: client,
		}
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case clientCreatedMessage:
		m.client = msg.client
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab", "shift+tab", "up", "down":
			s := msg.String()
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}
			cmds := make([]tea.Cmd, len(m.inputs))

			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				m.inputs[i].Blur()
			}

			if m.focusIndex >= len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = 0
			}

			return m, tea.Batch(cmds...)
		case "enter":
			m.client.SendMsg(context.Background(), &chat.Message{Author: m.inputs[0].Value(), Content: m.inputs[1].Value()})
			m.inputs[0].SetValue("")
			m.inputs[1].SetValue("")
		}
	}
	cmd := m.updateInputs(msg)
	return m, cmd
}

func (m model) View() string {
	s := ""

	s += m.messageList.View()

	s += "--------------------- \n"

	for _, input := range m.inputs {
		s += input.View() + "\n"
	}

	return s
}
