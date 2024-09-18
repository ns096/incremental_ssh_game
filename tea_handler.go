package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish/bubbletea"
	zone "github.com/lrstanley/bubblezone"
)

// You can wire any Bubble Tea model up to the middleware with a function that
// handles the incoming ssh.Session. Here we just grab the terminal info and
// pass it to the new model. You can also return tea.ProgramOptions (such as
// tea.WithAltScreen) on a session by session basis.
func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	// This should never fail, as we are using the activeterm middleware.
	pty, _, _ := s.Pty()

	// When running a Bubble Tea app over SSH, you shouldn't use the default
	// lipgloss.NewStyle function.
	// That function will use the color profile from the os.Stdin, which is the
	// server, not the client.
	// We provide a MakeRenderer function in the bubbletea middleware package,
	// so you can easily get the correct renderer for the current session, and
	// use it to create the styles.
	// The recommended way to use these styles is to then pass them down to
	// your Bubble Tea model.
	renderer := bubbletea.MakeRenderer(s)
	txtStyle := renderer.NewStyle().Foreground(lipgloss.Color("10"))
	quitStyle := renderer.NewStyle().Foreground(lipgloss.Color("8"))

	// bubble zone
	zone.NewGlobal()

	items := []list.Item{
		// an ID field has been added here, however it's not required. You could use
		// any text field as long as it's unique for the zone.
		item{id: "item_1", title: "Raspberry Pi’s", desc: "I have ’em all over my house"},
		item{id: "item_2", title: "Nutella", desc: "It's good on toast"},
		item{id: "item_3", title: "Bitter melon", desc: "It cools you down"},
		item{id: "item_4", title: "Nice socks", desc: "And by that I mean socks without holes"},
		item{id: "item_5", title: "Eight hours of sleep", desc: "I had this once"},
		item{id: "item_6", title: "Cats", desc: "Usually"},
		item{id: "item_7", title: "Plantasia, the album", desc: "My plants love it too"},
	}

	bg := "light"
	if renderer.HasDarkBackground() {
		bg = "dark"
	}

	m := model{
		term:      pty.Term,
		profile:   renderer.ColorProfile().Name(),
		width:     pty.Window.Width,
		height:    pty.Window.Height,
		bg:        bg,
		list:      list.New(items, list.NewDefaultDelegate(), 80, 10),
		txtStyle:  txtStyle,
		quitStyle: quitStyle,
	}
	m.list.Title = "Left click on an items title to select it"

	return m, []tea.ProgramOption{tea.WithAltScreen(), tea.WithMouseCellMotion()}
}

type item struct {
	id    string
	title string
	desc  string
}

func (i item) Title() string       { return zone.Mark(i.id, i.title) }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return zone.Mark(i.id, i.title) }

// Just a generic tea.Model to demo terminal information of ssh.
type model struct {
	term      string
	profile   string
	width     int
	height    int
	bg        string
	list      list.Model
	txtStyle  lipgloss.Style
	quitStyle lipgloss.Style
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}

	case tea.MouseMsg:
		if msg.Button == tea.MouseButtonWheelUp {
			m.list.CursorUp()
			return m, nil
		}

		if msg.Button == tea.MouseButtonWheelDown {
			m.list.CursorDown()
			return m, nil
		}

		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonLeft {
			for i, listItem := range m.list.VisibleItems() {
				v, _ := listItem.(item)
				// Check each item to see if it's in bounds.
				if zone.Get(v.id).InBounds(msg) {
					// If so, select it in the list.
					m.list.Select(i)
					break
				}
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	docStyle := lipgloss.NewStyle().Margin(1, 2)
	s := fmt.Sprintf("Your term is %s\nYour window size is %dx%d\nBackground: %s\nColor Profile: %s", m.term, m.width, m.height, m.bg, m.profile)
	return m.txtStyle.Render(s) + zone.Scan(docStyle.Render(m.list.View())) + "\n\n" + m.quitStyle.Render("Press 'q' to quit\n")
}
