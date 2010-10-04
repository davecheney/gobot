package main

import (
	"log"
	"bufio"
	"os"
	"net"
	"fmt"
)

type IRCWriter chan string 

type IRCReader chan string

type IRC struct {
 	Reader chan string
	Writer IRCWriter
}

func (w IRCWriter) Send(line string) {
 	w <- line
}

func (w IRCWriter) Printf(format string, a ...interface{}) {
	w <- fmt.Sprintf(format, a)
}

func (r IRCReader) ReadLine() string {
	return <- r
}

func newIRC(c net.Conn) (irc *IRC) {
	irc = &IRC{ Reader: make(chan string, 10), Writer: make(chan string, 10) }

	// start reader
	go func(r *bufio.Reader) {
		for {
			line, err := r.ReadString('\n')
			if err != nil {
				log.Exit("Unable to read from channel", err)
			}
			log.Stdoutf("Read: [%s]", line)
			irc.Reader <- line
		}
	}(bufio.NewReader(c))
	
	// start writer
	go func(w *bufio.Writer) {
		for {
			line := <- irc.Writer
			log.Stdout(line)
			_, err := w.WriteString(line)
			if err != nil {
				log.Exit("Unable to write to channel", err)
			}
			err = w.Flush()
			if err != nil {
				log.Exit("Unable to write to channel", err)
			}

			log.Stdoutf("Wrote: [%s]", line)
		}
	}(bufio.NewWriter(c))

	return irc
}

func connectTLS(host string, port int) (irc *IRC,err os.Error)  {
	conn, err := dialIRCTLS(host, port)
	if err != nil {
		return nil, err
	}
	return newIRC(conn), nil
}
