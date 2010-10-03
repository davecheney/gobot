package main

import (
	"flag"
	"fmt"
	
	// "net"
	"log"
	// "os"
	
)

var host *string = flag.String("host", "", "IRC Host")

func main() {
	flag.Parse()
	
	conn, err := dialIRCTLS(*host, 6669)
	if err != nil {
		log.Exit("Unable to dial IRC server", err)
	}
	
	buf := make([]byte, 4096)
	read, err := conn.Read(buf)
	fmt.Println("read", read, buf[0:read])
	if err != nil {
		log.Exit("Unable to read from IRC", err)
	}
}
