package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var timeout time.Duration

func init() {
	flag.DurationVar(&timeout, "timeout", 10*time.Second, "timeout connection")
}

func printError(err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
	os.Exit(1)
}

func receiveLoop(client TelnetClient, ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := client.Receive()
			if err != nil {
				fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
				os.Exit(1)
			}
		}
	}
}

func sendLoop(client TelnetClient, ticker *time.Ticker, done chan bool) {
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			err := client.Send()
			if err != nil {
				fmt.Fprintln(os.Stderr, "...EOF")
			}
		}
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		printError(errors.New("not found host or port"))
	}
	address := net.JoinHostPort(flag.Arg(0), flag.Arg(1))
	client := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	printError(err)
	defer client.Close()
	tickerSend := time.NewTicker(500 * time.Millisecond)
	doneSend := make(chan bool)
	tickerReceive := time.NewTicker(500 * time.Millisecond)
	doneReceive := make(chan bool)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGQUIT)
	fmt.Fprintln(os.Stderr, "...Connected to", address)
	go sendLoop(client, tickerSend, doneSend)
	go receiveLoop(client, tickerReceive, doneReceive)
	<-sigs
	tickerSend.Stop()
	doneSend <- true
	tickerReceive.Stop()
	doneReceive <- true
}
