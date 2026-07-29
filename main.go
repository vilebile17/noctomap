package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	tea "charm.land/bubbletea/v2"
)

type Theme struct {
	wallpaperPath      string
	colourSchemeSource string
	colourSchemeName   string
}

type model struct {
	themes []Theme
	cursor int
}

func initialModel() model {
	return model{
		themes: []Theme{
			{
				wallpaperPath:      "~/Pictures/Wallpapers/Arch Linux.jpg",
				colourSchemeSource: "builtin",
				colourSchemeName:   "Catppuccin",
			},
			{
				wallpaperPath:      "~/Pictures/Wallpapers/Dramatic Thunder.png",
				colourSchemeSource: "builtin",
				colourSchemeName:   "Nord",
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.themes)-1 {
				m.cursor++
			}
		case "enter", "space":
			themeCmd := exec.Command(
				"noctalia",
				"msg",
				"color-scheme-set",
				m.themes[m.cursor].colourSchemeSource,
				m.themes[m.cursor].colourSchemeName,
			)
			err := themeCmd.Run()
			if err != nil {
				log.Fatal(err)
			}

			wallpaperCmd := exec.Command(
				"noctalia",
				"msg",
				"wallpaper-set",
				m.themes[m.cursor].wallpaperPath,
			)
			err = wallpaperCmd.Run()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	s := "What should we buy at the market?\n\n"

	for i, theme := range m.themes {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		s += fmt.Sprintf("%s %s\n", cursor, theme.wallpaperPath)
	}

	s += "\nPress q to quit.\n"

	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
