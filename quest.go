package main

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/log"
)

type quest struct {
	name             string
	description      string
	goal             int
	current_progress int
	pictogram        string
	progress         progress.Model
}

func createQuest(name string, goal int) quest {
	quest := quest{
		name:        name,
		description: "",
		goal:        14,
		pictogram:   createGoblin(),
		progress:    progress.New(progress.WithDefaultGradient()),
	}
	return quest
}

type updateMsg struct {
	tick           time.Time
	extra_progress int
}

func updateQuest(progress int, quest quest) quest {
	quest.current_progress += progress

	if quest.progress.Percent() >= 1.0 {
		log.Info("Quest over")
	}
	return quest
}

func viewQuest(quest quest) string {
	progress_in_percent := float64(quest.current_progress) / float64(quest.goal)
	log.Info("update ", progress_in_percent)

	progress := quest.progress.ViewAs(progress_in_percent)
	questView := fmt.Sprintf("%s\n%s\n%s", progress, quest.name, quest.pictogram)
	log.Info(quest)
	return questView

}
