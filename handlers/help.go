package handlers

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"
	"io"
)

func HelpProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{

	return telsh.PromoteHandlerFunc(helpHandler)
}

func helpHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {

	oi.LongWriteString(stdout, "Available Commands:\r\n")
	oi.LongWriteString(stdout, "help  - display this help menur\r\n")
	oi.LongWriteString(stdout, "dance - dance\r\n")
	oi.LongWriteString(stdout, "five  - you can count right?\r\n")
	oi.LongWriteString(stdout, "\r \r\n")

	return nil
}

