package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	themes []Theme
	cursor int
}

func initialModel() model {
	t, err := getConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	return model{
		themes: t,
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
		case "ctrl+c", "q", "esc":
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
				m.themes[m.cursor].ColourSchemeSource,
				m.themes[m.cursor].ColourSchemeName,
			)
			err := themeCmd.Run()
			if err != nil {
				return m, tea.Println(fmt.Sprintf(
					"Unable to change to theme %v -> %v, make sure you typed it out correctly\n(watch out for capitals!!)\n",
					m.themes[m.cursor].ColourSchemeSource,
					m.themes[m.cursor].ColourSchemeName,
				))
			}

			wallpaperCmd := exec.Command(
				"noctalia",
				"msg",
				"wallpaper-set",
				m.themes[m.cursor].WallpaperPath,
			)
			err = wallpaperCmd.Run()
			if err != nil {
				return m, tea.Println(fmt.Sprintf(
					"Unable to change wallpaper to '%v' be sure that you spelled it out correctly\n",
					m.themes[m.cursor].WallpaperPath,
				))
			}
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	s := "Which wallpaper are you craving?\n\n"

	for i, theme := range m.themes {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}

		directories := strings.Split(theme.WallpaperPath, "/")
		s += fmt.Sprintf("%s \033[1m%s\033[0m - %s\n",
			cursor,
			directories[len(directories)-1],
			theme.ColourSchemeName,
		)
	}

	s += "\nArrow keys to move around, enter to select and q to quit\n"

	return tea.NewView(s)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
