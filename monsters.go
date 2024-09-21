package main

import "github.com/charmbracelet/log"

func createGoblin() string {

	body := green.Render("╔╩╗")
	head := light_green.Render(" ,│\n╔▀╬")
	goblin := head + "\n" + body
	log.Info(goblin)
	return goblin
}
