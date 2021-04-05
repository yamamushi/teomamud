package handlers

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"

	"io"
)

func FiveProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	return telsh.PromoteHandlerFunc(fiveHandler)
}

func fiveHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "The number FIVE looks like this: 5\r\n")

	return nil
}

