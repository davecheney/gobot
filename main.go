package main

import (
	"flag"
	"log"
)

var host *string = flag.String("host", "", "IRC Host")
var port *int = flag.Int("port", 6669, "IRC Port")
var user *string = flag.String("user", "", "USER")
var nick *string = flag.String("nick", "gobot", "NICK")
var pass *string = flag.String("pass", "", "PASS")
var join *string = flag.String("join", "", "JOIN")

func main() {
	flag.Parse()
	irc, err := connectTLS(*host, *port)
	if err != nil {
		log.Exit("Unable to read from IRC", err)
	}
	
	irc.Writer.Printf("PASS %s\r\nUSER %s foo foo :gobot\r\nNICK gobot\r\n", *pass, *user)
	irc.Writer.Printf("JOIN %s\r\n", *join)

	bot := &Bot{"#bacon"}
	
	for {
		bot.Accept( <- irc.Reader, irc.Writer)
	}
}
