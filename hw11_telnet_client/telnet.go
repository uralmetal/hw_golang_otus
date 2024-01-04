package main

import (
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type telnetClient struct {
	TelnetClient
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	conn    net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &telnetClient{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (client *telnetClient) Connect() (err error) {
	client.conn, err = net.DialTimeout("tcp", client.address, client.timeout)
	return
}

func (client *telnetClient) Send() (err error) {
	buffer := make([]byte, 1024)
	n, err := client.in.Read(buffer)
	if err != nil {
		return
	}
	_, err = client.conn.Write(buffer[:n])
	return
}

func (client *telnetClient) Receive() (err error) {
	buffer := make([]byte, 1024)
	n, err := client.conn.Read(buffer)
	if err != nil {
		return
	}
	_, err = client.out.Write(buffer[:n])
	return
}

func (client *telnetClient) Close() (err error) {
	return client.conn.Close()
}
