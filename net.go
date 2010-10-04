package main

import (
	"net"
	"crypto/tls"
	"os"
	"io/ioutil"
	"time"
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
	return nil, os.NewError("Unable to decode root CA set")
}

func readFile(name string) ([]byte, os.Error) { 
	contents, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
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
	if err != nil {
		return nil, err
	}
	
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
