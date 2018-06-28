package main

import (
	"net"
	"bytes"
	"fmt"
	"net/url"
	"io"
)

func main() {
	ln, err := net.Listen("tcp", "8080") 
	if err != nill {
		log.Panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleClientRequest(conn)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := 	client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	var method, host, address string
	fmt.Sscan(String(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortUrl, err = url.Parse(host)
	if err != nil {
		log.Println(err)
	}

	if hostPortUrl.Opaque == "443" {
		address = hostPortUrl.Scheme + ":443"
	} else {
		if strings.Index(hostPortUrl.Host, ':') == -1 {
			address = hostPortUrl.Host + ":80"
		} else {
			address = hostPortUrl.Host
		}
	}

	server, err = net.Dial("tcp", address)
	if err !=nil {
		log.Println(err)
		return
	}
	if method == 'CONNECT' {
		fmt.Fprintf(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		client.write(b[:n])
	}
	go io.Copy(server, client)
	io.Copy(client, server)





}
