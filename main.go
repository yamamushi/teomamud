package main

import (
	"github.com/reiver/go-oi"
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"

	"github.com/yamamushi/teomamud/handlers"

	"io"
)



func fiveHandler(stdin io.ReadCloser, stdout io.WriteCloser, stderr io.WriteCloser, args ...string) error {
	oi.LongWriteString(stdout, "The number FIVE looks like this: 5\r\n")

	return nil
}

func fiveProducer(ctx telnet.Context, name string, args ...string) telsh.Handler{
	return telsh.PromoteHandlerFunc(fiveHandler)
}



func main() {

	shellHandler := telsh.NewShellHandler()

	shellHandler.WelcomeMessage = `
 __          __ ______  _        _____   ____   __  __  ______ 
 \ \        / /|  ____|| |      / ____| / __ \ |  \/  ||  ____|
  \ \  /\  / / | |__   | |     | |     | |  | || \  / || |__   
   \ \/  \/ /  |  __|  | |     | |     | |  | || |\/| ||  __|  
    \  /\  /   | |____ | |____ | |____ | |__| || |  | || |____ 
     \/  \/    |______||______| \_____| \____/ |_|  |_||______|

`


	// Register the "five" command.
	commandName     := "five"
	commandProducer := telsh.ProducerFunc(fiveProducer)

	shellHandler.Register(commandName, commandProducer)


	// Register the "dance" command.
	commandName      = "dance"
	commandProducer  = telsh.ProducerFunc(handlers.DanceProducer)

	shellHandler.Register(commandName, commandProducer)


	shellHandler.Register("dance", telsh.ProducerFunc(handlers.DanceProducer))

	addr := ":420"
	if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
		panic(err)
	}
}