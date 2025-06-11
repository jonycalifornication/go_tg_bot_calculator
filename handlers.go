package main

import (
	"fmt"
	tele "gopkg.in/telebot.v4"
)

var waitingForReply = make(map[int64]bool)

func registerHandlers(b *tele.Bot) {
	b.Handle("/start", handleStart)
	b.Handle(tele.OnText, handleText)
}

func handleStart(c tele.Context) error {
	waitingForReply[c.Sender().ID] = true
	text := fmt.Sprintf(`Привет, %s! 👋
	Я — Telegram-калькулятор 🤖
	Я умею вычислять простые математические выражения.
	Поддерживаю операции: +, -, *, /
	Просто отправь мне выражение — и я посчитаю!
	Напиши stop, чтобы закончить сессию.`, c.Sender().FirstName)
	return c.Send(text)
}

func handleText(c tele.Context) error {
	userID := c.Sender().ID
	result := calculator(c.Text())
	if c.Text() == "stop" {
		waitingForReply[userID] = false
		return c.Send("Stopped")
	}
	return c.Send(result)
}
