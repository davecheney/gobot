package main

import (
	"log"
	"regexp"
)

var (
	R = regexp.MustCompile("^(:([a-z0-9\\.\\-@!]+) )?([a-zA-Z0-9]+)( (.*))?$")
)

type Bot struct {
	Channel string
}

type Message struct {
	Command string
	Params string
}

func (m *Message) Ping(out IRCWriter) {
	out.Printf("PONG %s\r\n", "foo")
}

func (bot *Bot) Process(message *Message, out chan string) {
	switch message.Command {
		case "PING": message.Ping(out)
		default: log.Stdoutf("%#v", message)
	}
}

func (bot *Bot) Accept(line string, out chan string) {
	command := R.FindStringSubmatch(line)
	if command != nil {
		log.Stdout("%#v", command)
		bot.Process(&Message{ 
			Command: command[3],
			Params: command[5],
		}, out)
	}
}
