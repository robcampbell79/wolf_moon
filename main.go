package main

import(
	"fmt"
	"strings"
	_ "io"
	_ "os"
	_ "bufio"

	"wolf_moon/chapters"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewportA	viewport.Model
	viewportB	viewport.Model
	textInput	textinput.Model
	mode		int
}

func initialModel() model {
	vp1 := viewport.New(300, 20)
	vp1.Style = lipgloss.NewStyle().
		Width(300).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("91")).
		Foreground(lipgloss.Color("163"))

		intro, md := chapters.GetChapter(1, 0)
	
	vp1.SetContent(intro)

	vp2 := viewport.New(100, 20)
	vp2.Style = lipgloss.NewStyle().
		Align(lipgloss.Left).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("91")).
		Foreground(lipgloss.Color("163"))

	vp2.SetContent("-----Item List-----\n\nEmpty")

	ti := textinput.New()
	ti.Placeholder = "Enter Your Actions:"
	ti.Focus()

	return model {
		viewportA: vp1,
		viewportB: vp2,
		textInput: ti,
		mode: md,
	}

}

func(m model) Init() tea.Cmd {
	return nil
}

func(m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// var(
	// 	tiCmd tea.Cmd
	// 	vpCmd tea.Cmd
	// )

	// m.textInput, tiCmd = m.textInput.Update(msg)
	// m.viewportA, vpCmd = m.viewportA.Update(msg)
	m.textInput, _ = m.textInput.Update(msg)
	m.viewportA, _ = m.viewportA.Update(msg)
	var vstr []string

	var chp int
	var stp int

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.textInput.Value() == "exit" || m.textInput.Value() == "quit" || m.textInput.Value() == "q" {
				return m, tea.Quit
			}

			if m.mode == 5 {
				chp = 1
				stp = stp + 1
				step, _ := chapters.GetChapter(chp, stp)
				vstr = append(vstr, step)
			}
			
			
			m.viewportA.SetContent(strings.Join(vstr, ""))

			m.textInput.Reset()
		}
	}

	return m ,nil
}

func(m model) View() string {

	views := []string{m.viewportA.View(), m.viewportB.View()}

	return lipgloss.JoinHorizontal(lipgloss.Top, views...) + "\n\n" + fmt.Sprintf("%s",m.textInput.View(),) + "\n\n"
	// return lipgloss.JoinHorizontal(lipgloss.Top, views...) + "\n\n"

	
}

func main() {

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
	}
}