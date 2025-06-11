package main

import tele "gopkg.in/telebot.v4"

var waitingForReply = make(map[int64]bool)

func registerHandlers(b *tele.Bot) {
	b.Handle("/start", handleStart)
	b.Handle(tele.OnText, handleText)
}

func handleStart(c tele.Context) error {
	waitingForReply[c.Sender().ID] = true
	return c.Send("Write your text:")
}

func handleText(c tele.Context) error {
	userID := c.Sender().ID
	result := calculator(c.Text())
	if c.Text() == "stop" {
		waitingForReply[userID] = false
		return c.Send("Stopped")
	}
	return c.Send("Your text: " + result)
}
