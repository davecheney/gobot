package main

import (
	"log"
	"bufio"
	"os"
	"crypto/tls"
)

type IRCChannel struct {
	reader *bufio.Reader
	writer *bufio.Writer
	Reader chan string
	Writer chan string
}

func (irc *IRCChannel) startReader() {
	for {
		line, err := irc.reader.ReadString('\n')
		if err != nil {
			log.Exit("Unable to read from channel", err)
		}
		irc.Reader <- line
	}
}

func (irc *IRCChannel) startWriter() {
	for {
		line := <- irc.Writer
		log.Stdout(line)
		count, err := irc.writer.WriteString(line)
		if err != nil {
			log.Exit("Unable to write to channel", err)
		}
		err = irc.writer.Flush()
		if err != nil {
			log.Exit("Unable to write to channel", err)
		}
		
		log.Stdout("wrote", count)
	}
}

func newIRCChannel(conn *tls.Conn) *IRCChannel {
	c := &IRCChannel{
		reader: bufio.NewReader(conn),
		writer: bufio.NewWriter(conn),
		Reader: make(chan string, 10),
		Writer: make(chan string, 10),
	}
	go c.startReader()
	go c.startWriter()
	return c
}

func connect(host string, port int) (*IRCChannel, os.Error) {
	conn, err := dialIRCTLS(host, port)
	if err != nil {
		return nil, err
	}
	return newIRCChannel(conn), nil
}
