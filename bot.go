package main

import (
	"log"
)

type Bot struct {
	Channel string
}

func (bot *Bot) Accept(line string, out chan string) {
	log.Stdout(line)
}
