package main

import (
	"fmt"
	"log"

	"github.com/Sagleft/uchatbot-engine"
)

const (
	configJSONPath = "config.json"
)

func main() {
	app := newSolution()

	err := checkErrors(
		app.parseConfig,
		app.initBot,
	)
	if err != nil {
		log.Fatalln(err)
	}

	app.printLaunched()
	app.runInBackground()
}

func newSolution() *solution {
	return &solution{}
}

func (app *solution) printLaunched() {
	fmt.Println("bot started")
}

func (app *solution) runInBackground() {
	forever := make(chan bool)
	// run in background
	<-forever
}

func (app *solution) initBot() error {
	var err error
	app.Engine, err = uchatbot.NewChatBot(uchatbot.ChatBotData{
		Client: &app.Config.Utopia,
		Chats:  app.Config.Chats,
		Callbacks: uchatbot.ChatBotCallbacks{
			OnContactMessage:        app.onContactMessage,
			OnChannelMessage:        app.onChannelMessage,
			OnPrivateChannelMessage: app.onPrivateChannelMessage,
		},
		UseErrorCallback: true,
		ErrorCallback:    app.onError,
	})
	return err
}

func (app *solution) onError(err error) {
	fmt.Println(err)
}