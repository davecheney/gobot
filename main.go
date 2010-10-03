package main

import (
	"flag"
	"fmt"
	
	// "net"
	"log"
	// "os"
	"bufio"
)

var host *string = flag.String("host", "", "IRC Host")

func main() {
	flag.Parse()
	
	conn, err := dialIRCTLS(*host, 6669)
	if err != nil {
		log.Exit("Unable to dial IRC server", err)
	}
	
	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')
	if err != nil {
		log.Exit("Unable to read from IRC", err)
	}
	fmt.Printf(line)
}
