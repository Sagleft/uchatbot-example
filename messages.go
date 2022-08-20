package main

import (
	"fmt"

	utopiago "github.com/Sagleft/utopialib-go"
)

func (app *solution) onContactMessage(m utopiago.InstantMessage) {
	fmt.Println("[CONTACT] " + m.Nick + ": " + m.Text)
}

func (app *solution) onChannelMessage(m utopiago.ChannelMessage) {
	fmt.Println("[CHANNEL] " + m.Nick + ": " + m.Text)
}

func (app *solution) onPrivateChannelMessage(m utopiago.ChannelMessage) {
	fmt.Println("[PRIVATE] " + m.Nick + ": " + m.Text)
}
