Title: Incremental ssh game
Date: 24-09-2024
#gamedev 
# Installation and running the game
- Install go from https://go.dev

`go run .`
`ssh -p 1338 user@localhost`

## Game Summary Pitch
This is an SSH app that allows any connection.
The session persists per user and continues when the user reconnects.
The user gets quests and clicks through them. Collecting resources and expanding his domain. 
1. Craft your armor and weapon
2. Slay some goblins
3. Get some followers
4. Conquer the land
5. Create your kingdom
6. Fight the *Evil Empire*
7. ???
8. Profit


## Inspiration
- Cookie Clicker
- Basic Fantasy stories
- Beautiful TUI 

# Technology
## Platform
- Any terminal that connects over SSH

## Development Software
- Language: Go
- Frameworks provided by https://charm.sh: BubbleTea, Wish, Lipgloss
- Database: SqlLite for session persistence

## Architecture
- Elm Architecture
- Data driven quests
- SSH server
- Single global state
- SqlLite DB

## Genre
- Incremental Game
- Idle Games

## Target Audience
- Terminal developers
- Fans of Incremental & Idle games

# Concept
## Gameplay Overview
- The player clicks through quests and unlocks upgrades that allow him to gather even more resources
- The Player continues to click through personal quest lines
- Finishing Quests gives additional resources
- Player can select from different questlines
- The unlocked questlines escalate in scale
- Once unlocked the other quest boxes progress by themselves with occasional weakpoints to progress faster

## Primary Mechanics
- Fast Clicking 
- Buy upgrades
- Idle time
- Quick time events
    - 'Golden Cookies'
    - Quest image shows a weak point that allows for faster progress
- Player can gather experience, gold, materials


## Secondary Mechanics
- Questlines
- Dialogue box
- Resource gathering

# Art
## Theme
- Fantasy
- ASCII Art with colors
## Design
https://kirill-live.itch.io/ascii-art-paint

# Game Experience
## UI
Simple clickable boxes that each represent a running quest. With each quest escalating in scale


## Controls
- Mouse and Keyboard
