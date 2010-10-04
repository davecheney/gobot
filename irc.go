package main

import (
	"log"
	"bufio"
	"os"
)

type IRCChannel struct {
	Reader chan string
	Writer chan string
}

func makeReaderChannel(reader *bufio.Reader) chan string {
	c := make(chan string)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				log.Exit("Unable to read from channel", err)
			}
			c <- line
		}
	}()
	return c
}

func makeWriterChannel(writer *bufio.Writer) chan string {
	c := make(chan string, 10)
	go func() {
		for {
			line := <- c
			log.Stdout(line)
			count, err := writer.WriteString(line)
			if err != nil {
				log.Exit("Unable to write to channel", err)
			}
			err = writer.Flush()
			if err != nil {
				log.Exit("Unable to write to channel", err)
			}
			
			log.Stdout("wrote", count)
		}
	}()
	return c
}

func connect(host string, port int) (*IRCChannel, os.Error) {
	conn, err := dialIRCTLS(host, port)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	
	return &IRCChannel{makeReaderChannel(reader), makeWriterChannel(writer)}, nil
}
