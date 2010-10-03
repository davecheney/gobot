package main

import (
	"net"
	"crypto/tls"
	"os"
	"io"
	"bufio"
	"time"
	// "fmt"
)

func dialIRC(host string, port int) (*net.TCPConn, os.Error) {
	remote, err := net.ResolveIPAddr(host)
	if err != nil {
		return nil, err
	}
	
	return net.DialTCP("tcp4", nil, &net.TCPAddr{ remote.IP, port })
}

func loadRootCA(file string) (*tls.CASet, os.Error) {
	pemBytes, err := readFile(file)
	if err != nil {
		return nil, err
	}

	caset := tls.NewCASet()
	if caset.SetFromPEM(pemBytes) {
		return caset, nil
	}
	return nil, os.NewError("Unable to decode root ca set")
}

func readFile(name string) ([]byte, os.Error) { 
       file, err := os.Open(name, os.O_RDONLY, 0); 
       if err != nil { 
       return nil, err; 
       } 
       stat, err := file.Stat(); 
       if err != nil { 
       return nil, err; 
       } 
       contents := make([]byte, stat.Size); 
       _, err = io.ReadFull(file, contents); 
       if err != nil { 
       return nil, err; 
       } 
       file.Close(); 
       return contents, nil; 
}

func newConfig() (*tls.Config, os.Error) {
	rootca, err := loadRootCA("ca-certificates.crt")
	if err != nil {
		return nil, err
	}
	
	urandom, err := os.Open("/dev/urandom", os.O_RDONLY, 0)
	if err != nil { 
	   return nil, err
	}
	
	return &tls.Config{
		Rand: urandom,
		Time: time.Seconds,
		RootCAs: rootca,
	}, nil
}

func dialIRCTLS(host string, port int) (c *tls.Conn, err os.Error) {
	remote, err := dialIRC(host, port)

	config, err := newConfig()
	if err != nil {
		return nil, err
	}
	
	c = tls.Client(remote, config)
    err = c.Handshake()
    if err == nil {
        return c, nil
    }
	c.Close()
    return nil, err
}

func makeReaderChannel(reader *bufio.Reader) chan string {
	c := make(chan string)
	go func() {
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				return
			}
			c <- line
		}
	}()
	return c
}

func connect(host string, port int) (chan string, chan string, os.Error) {
	conn, err := dialIRCTLS(host, port)
	if err != nil {
		return nil, nil, err
	}
	reader := bufio.NewReader(conn)
	// writer := bufio.NewWriter(conn)
	
	return makeReaderChannel(reader), nil, nil
}
