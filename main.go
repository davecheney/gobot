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



func main() {
	flag.Parse()
	r, _, err := connect(*host, 6669)
	if err != nil {
		log.Exit("Unable to read from IRC", err)
	}
	
	for {
		line := <- r
		fmt.Printf(line)
	}

}
