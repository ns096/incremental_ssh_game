package main

import "github.com/charmbracelet/lipgloss"

var (
	questStyle = lipgloss.NewStyle().
			Width(30).
			Height(10).
			Align(lipgloss.Center, lipgloss.Top).
			BorderStyle(lipgloss.NormalBorder())

	bold        = lipgloss.NewStyle().Bold(true)
	green       = lipgloss.NewStyle().Foreground(lipgloss.Color("#4bb83e")).Inherit(bold)
	light_green = lipgloss.NewStyle().Foreground(lipgloss.Color("#3fec45")).Inherit(bold)
)
