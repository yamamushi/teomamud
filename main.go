package main

import (
	"github.com/reiver/go-telnet"
	"github.com/reiver/go-telnet/telsh"

	"github.com/yamamushi/teomamud/handlers"
)


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
	shellHandler.Register("five", telsh.ProducerFunc(handlers.FiveProducer))

	// Register the "dance" command.
	shellHandler.Register("dance", telsh.ProducerFunc(handlers.DanceProducer))

	// Register the "help" command.
	shellHandler.Register("help",telsh.ProducerFunc(handlers.HelpProducer))


	addr := ":420"
	if err := telnet.ListenAndServe(addr, shellHandler); nil != err {
		panic(err)
	}
}