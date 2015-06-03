package main

import (
	"fmt"
	"github.com/davidhampgonsalves/slackline/conf"
	"github.com/davidhampgonsalves/slackline/slack"
	"os"
)

/*

flags
 -c = channel names (slack channel names are single word seperated by comma)
 -t = set token
 --save = persist flags to config
 -i = init, ask for channels and token
handle piping input?

TODO:
  if no cred or channels, display link to get token from slack
    - allow user to enter token which will be saved to .slk-line encrypted with mac, cpu, etc as key
  init walkthrough
  message format
*/

func main() {
	config, err := conf.Load()
	fmt.Println(*config)

	if err != nil {
		fmt.Print(err)
		os.Exit(0)
	}

	err = slack.JoinChannel(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	err = slack.PostMessage(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	if config.Save {
		conf.Save(config)
	}
}
