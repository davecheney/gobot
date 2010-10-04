package main

import (
	"flag"
	"fmt"
	
	// "net"
	"log"
	// "os"
	// "bufio"
)

var host *string = flag.String("host", "", "IRC Host")
var user *string = flag.String("user", "", "USER")
var nick *string = flag.String("nick", "gobot", "NICK")
var pass *string = flag.String("pass", "", "PASS")


func main() {
	flag.Parse()
	irc, err := connect(*host, 6669)
	if err != nil {
		log.Exit("Unable to read from IRC", err)
	}
	
	irc.Writer <- "PASS "+*pass+"\r\nUSER "+*user+" foo foo :gobot\r\nNICK gobot\r\n"
	
	for {
		line := <- irc.Reader
		fmt.Printf(line)
	}

}
